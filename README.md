# Калькулятор выражений с асинхронной обработкой

Этот проект представляет собой веб-сервис для вычисления математических выражений с возможностью регистрации пользователей и хранением истории вычислений.

## Особенности

- Асинхронная обработка выражений с использованием архитектуры Orchestrator-Agent
- Аутентификация пользователей (JWT)
- Хранение истории вычислений
- Валидация входных выражений
- Фронтенд на React/Vite

## Технологический стек

**Бэкенд:**
- Go (1.16+)
- SQLite3
- gRPC

**Фронтенд:**
- React
- Vite
- Node.js

## Установка и запуск

### Предварительные требования

1. Установите [Go](https://golang.org/dl/) (версия 1.16 или выше)
2. Установите [SQLite3](https://www.sqlite.org/download.html)
3. Убедитесь, что CGO включен (`export CGO_ENABLED=1`)
4. Установите GCC компилятор
5. Для фронтенда установите [Node.js](https://nodejs.org/en/download)

### Запуск сервиса

1. Клонируйте репозиторий:

```bash
git clone https://github.com/Killered672/LastModule
cd LastModule
```

2. Запустите Orchestrator:

```bash
export ORCHESTRATOR_URL=localhost:50051
export TIME_ADDITION_MS=200
export TIME_SUBTRACTION_MS=200
export TIME_MULTIPLICATIONS_MS=300
export TIME_DIVISIONS_MS=400

go run cmd/orchestrator.start/main.go
```

3. В новом терминале запустите Agent:

```bash
export COMPUTING_POWER=4
export ORCHESTRATOR_URL=localhost:50051

go run cmd/agent.start/main.go
```

4. (Опционально) Запустите фронтенд:

```bash
cd frontend
npm install
npm run dev
```

Фронтенд будет доступен по адресу: http://localhost:5173

## Использование API

### Регистрация пользователя

```bash
curl -X POST http://localhost:8080/api/v1/register \
  -H "Content-Type: application/json" \
  -d '{"login":"user1","password":"password123"}'
```

### Аутентификация

```bash
curl -X POST http://localhost:8080/api/v1/login \
  -H "Content-Type: application/json" \
  -d '{"login":"user1","password":"password123"}'
```

### Отправка выражения на вычисление

```bash
curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer YOUR_JWT_TOKEN' \
--data '{"expression": "2+2*2"}'
```

Ответ:
```json
{
  "id": "task-id"
}
```

### Проверка статуса выражения

```bash
curl --location 'http://localhost:8080/api/v1/expressions' \
--header 'Authorization: Bearer YOUR_JWT_TOKEN'
```

Ответ:
```json
{
  "expressions": [
    {
      "id": "1",
      "expression": "2*2+2",
      "status": "completed",
      "result": 6
    }
  ]
}
```

## Примеры ошибок

- Пользователь уже существует:
```json
{"error":"User already exists"}
```

- Невалидное выражение:
```json
{"error":"expected number at position 2"}
```

- Деление на ноль:
```json
{"error":"division by zero"}
```

## Тестирование

Для запуска тестов:

```bash
go test ./internal/agent/agent_calculation_test.go
```

Или интеграционные тесты:

```bash
go test ./cmd/internal_test.go
```

Тесты хранилища:

```bash
go test ./internal/storage/storage_test.go
```

## Контакты

По вопросам обращайтесь: [Telegram](https://t.me/Killered_656)
