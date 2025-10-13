package client_cmd

import (
	"bufio"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"golang.org/x/crypto/curve25519"
)

func newAddCommand() *cobra.Command {
	var pubKeyServer string
	var clientIP string

	cmd := &cobra.Command{
		Use:   "new",
		Short: "Add new client",
		RunE: func(cmd *cobra.Command, args []string) error {
			if pubKeyServer == "" {
				stat, err := os.Stdin.Stat()
				if err != nil {
					return err
				}
				// Проверяем, что stdin не из tty, а из pipe
				if (stat.Mode() & os.ModeCharDevice) == 0 {
					reader := bufio.NewReader(os.Stdin)
					input, err := reader.ReadString('\n')
					if err != nil {
						return err
					}
					pubKeyServer = input
				} else {
					return fmt.Errorf("public key must be provided either via --pub_key flag or piped input")
				}
			}

			// Удалим лишние пробелы и переносы строк у ключа
			pubKeyServer = strings.TrimSpace(pubKeyServer)

			if len(args) != 1 {
				cmd.Println("give name client")
				return nil
			}

			name := args[0]

			if clientIP == "" {
				return fmt.Errorf("IP client must be provided via --client_ip")
			}

			privateClientKey, publicClientKey, err := GenerateWireGuardKeyPair()
			if err != nil {
				fmt.Println("Error:", err)
				return err
			}

			ip, err := GetExternalIP()
			if err != nil {
				return err
			}
			fmt.Println("External IP:", ip)

			err = CreateWireGuardConfig(name, ip, clientIP, privateClientKey, pubKeyServer)
			if err != nil {
				return err
			}

			err = CreateWireGuardPartsOfServerConfig("server_"+name, publicClientKey, clientIP)

			cmd.Printf("files for client %s created.\n", name)

			return nil
		},
	}

	cmd.Flags().StringVar(&pubKeyServer, "pub_key", "", "Public key of server")
	cmd.Flags().StringVar(&clientIP, "client_ip", "", "IP address of client")

	return cmd
}

func GenerateWireGuardKeyPair() (privateKey string, publicKey string, err error) {
	// Генерация 32-байтового приватного ключа
	priv := make([]byte, 32)
	_, err = rand.Read(priv)
	if err != nil {
		return "", "", err
	}

	// Форматирование приватного ключа по стандарту WireGuard
	// (исправление битов согласно RFC7748)
	priv[0] &= 248
	priv[31] &= 127
	priv[31] |= 64

	// Вычисление публичного ключа
	pub, err := curve25519.X25519(priv, curve25519.Basepoint)
	if err != nil {
		return "", "", err
	}

	// Кодирование ключей в base64 для удобства использования
	privateKey = base64.StdEncoding.EncodeToString(priv)
	publicKey = base64.StdEncoding.EncodeToString(pub)
	return privateKey, publicKey, nil
}

type IPResponse struct {
	Query string `json:"query"`
}

func GetExternalIP() (string, error) {
	resp, err := http.Get("http://ip-api.com/json/")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var ipResp IPResponse
	err = json.Unmarshal(body, &ipResp)
	if err != nil {
		return "", err
	}

	return ipResp.Query, nil
}

func CreateWireGuardConfig(fileName, serverIP, clientIP, clientPrivateKey, serverPublicKey string) error {
	const configTemplate = `[Interface]
PrivateKey = %s
Address = %s/32
DNS = 8.8.8.8

[Peer]
PublicKey = %s
Endpoint = %s:51830
AllowedIPs = 0.0.0.0/0
PersistentKeepalive = 20
`

	// Формируем содержимое конфигурации
	configContent := fmt.Sprintf(configTemplate, clientPrivateKey, clientIP, serverPublicKey, serverIP)

	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(configContent)
	if err != nil {
		return err
	}

	return nil
}

func CreateWireGuardPartsOfServerConfig(fileName, publicKey, clientIP string) error {
	const configTemplate = `

[Peer]
PublicKey = %s
AllowedIPs = %s/32`

	configContent := fmt.Sprintf(configTemplate, publicKey, clientIP)
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write([]byte(configContent))
	if err != nil {
		return err
	}

	return nil
}
