# � Expression Calculator Service

![Go](https://img.shields.io/badge/Go-1.16+-blue)
![SQLite](https://img.shields.io/badge/SQLite-3-lightgrey)
![Node.js](https://img.shields.io/badge/Node.js-18+-green)

Веб-сервис для вычисления математических выражений с аутентификацией пользователей и хранением истории вычислений.

## 🛠 Требования

Перед началом работы убедитесь, что у вас установлено:

- Go (версия 1.16 или выше)
- SQLite3
- GCC компилятор (для CGO)
- Node.js (для фронтенда)

### Установка зависимостей

**Windows:**
```powershell
winget install --id=SQLite.SQLite -e
```

**Linux/macOS:**
```bash
# Для SQLite (если не установлен)
sudo apt-get install sqlite3 libsqlite3-dev
```

## ⚙️ Установка

1. Клонируйте репозиторий:
```bash
git clone https://github.com/9doo/lastask
cd LastModule
```

2. Включите CGO (если выключен):
```bash
export CGO_ENABLED=1
```

## 🚀 Запуск

### 1. Запуск Orchestrator

В первом терминале:
```bash
export ORCHESTRATOR_URL=localhost:50051
export TIME_ADDITION_MS=200
export TIME_SUBTRACTION_MS=200
export TIME_MULTIPLICATIONS_MS=300
export TIME_DIVISIONS_MS=400

go run cmd/orchestrator.start/main.go
```

Ожидаемый вывод:
```
Starting Orchestrator on port 8080
Starting gRPC server on port 50051
Starting HTTP server on port 8080
```

### 2. Запуск Agent

Во втором терминале:
```bash
cd LastModule
export COMPUTING_POWER=4
export ORCHESTRATOR_URL=localhost:50051

go run cmd/agent.start/main.go
```

Ожидаемый вывод:
```
Starting Agent...
Starting worker 0
Starting worker 1
Starting worker 2
Starting worker 3
Worker X: error getting task: rpc error: code = Unknown desc = not found
```

*Примечание: Ошибки "not found" нормальны при отсутствии активных задач.*

## 🔐 Аутентификация

### Регистрация пользователя
```bash
curl -X POST http://localhost:8080/api/v1/register \
  -H "Content-Type: application/json" \
  -d '{"login":"user1","password":"password123"}'
```

### Вход пользователя
```bash
curl -X POST http://localhost:8080/api/v1/login \
  -H "Content-Type: application/json" \
  -d '{"login":"user1","password":"password123"}'
```

*Сохраните полученный JWT токен для последующих запросов.*

## 🧮 Использование API

### Отправка выражения на вычисление
```bash
curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer YOUR_JWT_TOKEN' \
--data '{"expression": "2+2*2"}'
```

Ответ:
```json
{"id": "..."}
```

### Проверка статуса вычислений
```bash
curl --location 'http://localhost:8080/api/v1/expressions' \
--header 'Authorization: Bearer YOUR_JWT_TOKEN'
```

Пример ответа:
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

## 💻 Фронтенд

1. Перейдите в папку фронтенда:
```bash
cd lastmodule/frontend
```

2. Установите зависимости и запустите:
```bash
npm install
npm run dev
```

Фронтенд будет доступен по адресу: [http://localhost:5173](http://localhost:5173)

*Примечание: Бекенд должен быть запущен для работы фронтенда.*

## ❌ Ошибки

### 1. Ошибка регистрации (409 Conflict)
**Когда возникает**: При попытке зарегистрировать пользователя, который уже существует

**Пример запроса**:
```bash
curl -X POST http://localhost:8080/api/v1/register \
  -H "Content-Type: application/json" \
  -d '{"login":"user1","password":"password123"}'
```

**Ответ**:
```json
{
  "error": "User already exists"
}
```

### 2. Ошибка 404 (Not Found)
**Когда возникает**: При обращении к несуществующему endpoint

**Пример запроса**:
```bash
curl --location 'http://localhost:8080/api/v1/nonexistent'
```

**Ответ**:
```json
{
  "error": "API Not Found"
}
```

### 3. Ошибка валидации выражения (422 Unprocessable Entity)
**Когда возникает**: При отправке выражения с синтаксической ошибкой

**Пример запроса**:
```bash
curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer YOUR_JWT_TOKEN' \
--data '{"expression": "2+a"}'
```

**Ответ**:
```json
{
  "error": "expected number at position 2"
}
```

### 4. Ошибка недопустимого токена (400 Bad Request)
**Когда возникает**: При отправке выражения с недопустимыми символами

**Пример запроса**:
```bash
curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer YOUR_JWT_TOKEN' \
--data '{"expression": "2\0"}'
```

**Ответ**:
```json
{
  "error": "Invalid token"
}
```

### 5. Ошибка деления на ноль (500 Internal Server Error)
**Когда возникает**: При попытке выполнить деление на ноль

**Пример запроса**:
```bash
curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer YOUR_JWT_TOKEN' \
--data '{"expression": "2/0"}'
```

**Ответ в логах агента**:
```
Worker: error computing task: division by zero
```

**Ответ API**:
```json
{
  "id": "task-id-123"
}
```

### 6. Ошибка аутентификации (401 Unauthorized)
**Когда возникает**: При запросе без токена или с невалидным токеном

**Пример запроса**:
```bash
curl --location 'http://localhost:8080/api/v1/expressions'
```

**Ответ**:
```json
{
  "error": "Authorization header is required"
}
```

### 7. Ошибка неверных учетных данных (401 Unauthorized)
**Когда возникает**: При вводе неверного логина/пароля

**Пример запроса**:
```bash
curl -X POST http://localhost:8080/api/v1/login \
  -H "Content-Type: application/json" \
  -d '{"login":"wrong","password":"wrong"}'
```

**Ответ**:
```json
{
  "error": "Invalid credentials"
}
```
## 🧪 Тестирование

Запуск unit-тестов:
```bash
go test ./internal/agent/agent_calculation_test.go
```

Интеграционные тесты:
```bash
go test ./cmd/internal_test.go
```

Тесты хранилища:
```bash
go test ./internal/storage/storage_test.go
```

Успешный вывод:
```
ok   calc_service/internal/evaluator 0.001s
```

## 📌 Примеры

### Успешный запрос
```bash
curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer YOUR_JWT_TOKEN' \
--data '{"expression": "2+2*2"}'
```
*Примечание: В тестах агента может возникать конфликт с ErrDivisionByZero из-за дублирования в коде.*
