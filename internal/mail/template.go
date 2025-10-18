package mail

import (
	"bytes"
	"fmt"
	"html/template"
	"time"
)

type mailData struct {
	Date string
}

var mailTemplate = `Привет, это Виктор Т.

Это автоматически сгенерированное письмо. По твоей просьбе сконфигурированы конфиги для vpn.

Краткая справка по установке конфигурации WireGuard для клиента:
Установите WireGuard клиент на ваше устройство.

Для Windows

1. Скачайте официальный клиент WireGuard для Windows с сайта https://www.wireguard.com/install/ и установите его.

2. Запустите приложение WireGuard, нажмите "Add Tunnel" → "Add empty tunnel" для создания нового подключения.

3. В открывшемся окне вставьте содержимое конфигурационного файла WireGuard (например, peer1.conf), который был предварительно подготовлен с приватным ключом клиента и настройками сервера.

4. Назовите туннель (например, wg), нажмите "Save", затем "Activate" для подключения.

5. Для отключения можно нажать "Deactivate".

Опционально настройте брандмауэр Windows, разрешив UDP-трафик на используемом порту WireGuard (обычно 51820) для корректной работы VPN.

Для Apple (macOS и iOS)

1. Установите приложение WireGuard из Mac App Store для macOS или из App Store для iOS.

2. Откройте приложение и выберите "Add a tunnel" → "Create from file or archive", затем импортируйте конфигурационный файл WireGuard.

3. Назначьте имя туннелю и сохраните.

4. Включите туннель для подключения.

Рекомендуемая документация и ссылки
Официальный сайт с загрузками и инструкциями: https://www.wireguard.com/install/

Дата: {{.Date}}
`

func templateMail(name string) (bytes.Buffer, error) {
	date := time.Now().UTC()

	data := mailData{
		Date: date.Format("02.01.2006 15:04"),
	}

	var body bytes.Buffer
	t := template.Must(template.New("email").Parse(mailTemplate))
	err := t.Execute(&body, data)
	if err != nil {
		return bytes.Buffer{}, fmt.Errorf("Create mail template %w", err)
	}

	return body, nil
}
