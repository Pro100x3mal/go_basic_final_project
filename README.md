# 📝 Todo List Scheduler

## О проекте

Todo List Scheduler - это веб-приложение для управления задачами написанное на GO, созданное в рамках финального проекта по курсу "Go-разработчик с нуля" от Яндекс Практикума.  
Приложение позволяет создавать, просматривать, редактировать и удалять задачи, включая поддержку одноразовых и периодических задач.

## ✅ Выполненные задания

- [x] Реализована авторизация по паролю с выдачей JWT
- [x] Хранение задач в SQLite-файле
- [x] Поддержка повторяющихся задач по годам, дням, дням недели (без возможности указания по дням месяца!) и автоматический расчёт даты следующего выполнения
- [x] Поиск задач по ключевому слову и дате
- [x] Dockerfile для запуска проекта в контейнере

## 🛠 Дополнительные особенности

- Использован архитектурный стиль handler-service-repository
- Применена концепция dependency injection для разделения слоев и изоляции бизнес-логики
- Пакет chi используется для маршрутизации и подключения middleware

## 🚀 Запуск проекта локально

### ⚙️ Переменные окружения

Можно указать через `.env` или при запуске:

```
TODO_PORT=7540 # порт HTTP сервера
TODO_PASSWORD=admin # пароль для аутентификации с помощью JWT
TODO_JWT_SECRET=your-secret # секрет для генерации и валидации подписи JWT
TODO_DBFILE=scheduler.db # путь к файлу базы данных
```
### ⚙️ Флаги командной строки

```
-p # порт HTTP сервера
```

### ▶️ Команда запуска из исходного кода

> ⚠️ **Важно:** По умолчанию сервер запускается с установленным паролем `admin`.  
> Если вы измените пароль через переменную окружения `TODO_PASSWORD`, необходимо также обновить значение переменной `Token` в тестах в файле `./tests/settings.go`, чтобы тесты успешно проходили.

```bash
go run ./cmd/server
```

По умолчанию сервер запускается по адресу:

```
http://0.0.0.0:7540
```

## 🧪 Запуск тестов

Перед запуском тестов укажите настройки в `tests/settings.go`:

```go
var Port = 7540
var DBFile = "../scheduler.db"
var FullNextDate = false
var Search = true
var Token = `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTI3MTI0NzUsImhhc2giOiI4YzY5NzZlNWI1NDEwNDE1YmRlOTA4YmQ0ZGVlMTVkZmIxNjdhOWM4NzNmYzRiYjhhODFmNmYyYWI0NDhhOTE4In0.Ewcoh7c5Hb6mQcuuSjHz76DCfZU7rX7TBW9GGluOn8U`

```

Затем выполните:

```bash
go test ./tests
```

## 🐳 Сборка и запуск с Docker

### Сборка образа:

```bash
docker build -t go-todo-app .
```

### Запуск контейнера:

```bash
docker run -p 7540:7540 \
  -e TODO_PORT=7540 \
  -e TODO_PASSWORD=admin \
  -e TODO_JWT_SECRET=very-secret-key \
  -e TODO_DBFILE=scheduler.db \
  --rm go-todo-app
```

Откройте в браузере:

```
http://localhost:7540
```