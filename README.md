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

# Примеры запросов и возможных ошибок

## 1. Ошибки аутентификации

### Неправильный пароль при входе
**Запрос:**
```bash
curl -X POST http://localhost:8080/api/v1/login \
  -H "Content-Type: application/json" \
  -d '{"login":"user1","password":"wrongpassword"}'
```

**Ответ:**
```json
{
  "error": "invalid credentials"
}
```

### Отсутствие токена при запросе вычисления
**Запрос:**
```bash
curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{"expression": "2+2"}'
```

**Ответ:**
```json
{
  "error": "authorization header is required"
}
```

## 2. Ошибки валидации выражений

### Недопустимый символ в выражении
**Запрос:**
```bash
curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer YOUR_JWT_TOKEN' \
--data '{"expression": "2@2"}'
```

**Ответ:**
```json
{
  "error": "Invalid token '@' at position 1"
}
```

### Неполное выражение
**Запрос:**
```bash
curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer YOUR_JWT_TOKEN' \
--data '{"expression": "2+"}'
```

**Ответ:**
```json
{
  "error": "unexpected end of expression"
}
```

### Деление на ноль
**Запрос:**
```bash
curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer YOUR_JWT_TOKEN' \
--data '{"expression": "1/0"}'
```

**Ответ (в логах агента):**
```
Worker: error computing task: division by zero
```

## 3. Ошибки работы с задачами

### Запрос несуществующей задачи
**Запрос:**
```bash
curl --location 'http://localhost:8080/api/v1/expressions/999' \
--header 'Authorization: Bearer YOUR_JWT_TOKEN'
```

**Ответ:**
```json
{
  "error": "expression not found"
}
```

### Неверный формат ID задачи
**Запрос:**
```bash
curl --location 'http://localhost:8080/api/v1/expressions/abc' \
--header 'Authorization: Bearer YOUR_JWT_TOKEN'
```

**Ответ:**
```json
{
  "error": "invalid expression ID"
}
```

## 4. Ошибки регистрации

### Попытка повторной регистрации
**Запрос:**
```bash
curl -X POST http://localhost:8080/api/v1/register \
  -H "Content-Type: application/json" \
  -d '{"login":"user1","password":"password123"}'
```

**Ответ (если пользователь уже существует):**
```json
{
  "error": "User already exists"
}
```

### Неполные данные при регистрации
**Запрос (без пароля):**
```bash
curl -X POST http://localhost:8080/api/v1/register \
  -H "Content-Type: application/json" \
  -d '{"login":"newuser"}'
```

**Ответ:**
```json
{
  "error": "login and password are required"
}
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

