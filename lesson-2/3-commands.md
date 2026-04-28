Модули
```bash
go mod init <REPOSITORY_URL> # URL без http(s), например: gitlab.golang-school.ru/username/my-app
go mod init <REPOSITORY_URL> # Создаёт модуль с указанной версией
go mod tidy                  # Удаляет неиспользуемые зависимости из go.mod и go.sum
go mod download              # Загружает зависимости в кэш
go mod vendor                # Создаёт папку vendor с зависимостями
```

Установка бинарника в ~/go/bin
```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

Установка зависимостей
```bash
go get github.com/google/uuid           # Установка модуля
go get -u github.com/google/uuid        # Установка последней версии
go get github.com/google/uuid@v1.6.0    # Установка конкретной версии
```

Сборка приложения
```bash
go build -o my-app ./cmd/app  # Сборка приложения в бинарник с названием my-app
```

Запуск приложения
```bash
go run .                     # Ищет в текущем каталоге файл с функцией main
go run ./cmd/app/main.go     # Запуск с указанием конкретного файла
go run ./cmd/app             # Чаще делают так

go run -mod=vendor .         # Запуск с использованием зависимостей из папки vendor
```

Docker
```bash
docker build -t my-app:v0.1.0 .
docker run -p 8080:8080 my-app:v0.1.0
```

Тесты
```bash
go test ./...                # Запуск всех тестов в проекте
go test -v ./...             # С подробным выводом
go test -cover ./...         # С покрытием
```