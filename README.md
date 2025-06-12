# kubehealth

🩺 **kubehealth** — это минималистичный CLI-инструмент на Go для проверки состояния подов в Kubernetes-кластере.

## 🚀 Возможности

- Сканирует все namespace'ы в кластере
- Выводит общее количество подов в статусах:
  - `Running`
  - `Pending`
  - `Failed`
- Работает через локальный kubeconfig (`~/.kube/config`)

## 📦 Установка и запуск

```bash
git clone https://github.com/Saturnnnnnnnnnnn/kubehealth_project.git
cd kubehealth_project
go mod tidy
go run main.go check
```

Пример вывода:

```
Namespace: production
Running: 12, Pending: 1, Failed: 0

Namespace: dev
Running: 5, Pending: 0, Failed: 1

Namespace: kube-system
Running: 8, Pending: 0, Failed: 0
```

## 🧰 Зависимости

- [Go](https://golang.org/) >= 1.24.3
- [client-go](https://github.com/kubernetes/client-go)
- [Cobra](https://github.com/spf13/cobra)

## 📂 Структура проекта

```
kubehealth/
├── cmd/
│   └── check.go      # Логика команды `check`
├── main.go           # Точка входа и конфигурация Cobra
├── go.mod            # Модули и зависимости
└── README.md         # Этот файл
```

## 📌 Примечания

- Убедитесь, что у вас настроен kubeconfig по пути `~/.kube/config`.
- Проект можно расширить логикой для Telegram, JSON-вывода или HTTP API.

## 🧑‍💻 Автор

**Shermatov Khamza**  
Telegram: [@AllSaturn](https://t.me/AllSaturn)  
Email: hamza.shermatov0504@gmail.com
GitHub: https://github.com/Saturnnnnnnnnnnn/kubehealth
