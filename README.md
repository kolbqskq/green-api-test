# Green API Test

## Описание

Go-клиент для работы с Green API и набор интеграционных тестов для методов:

* sendMessage
* getChatHistory
* getStateInstance

---

## Инструкция по запуску

### 1. Установка Go

Требуется Go 1.26+

Linux:

```bash
sudo apt install golang
```

macOS:

```bash
brew install go
```

Windows:
Скачать и установить с официального сайта:
```
https://go.dev/doc/install
```

Проверка установки:

```bash
go version
```

---

### 2. Клонирование репозитория

```bash
git clone https://github.com/kolbqskq/green-api-test
cd green-api-test
```

---

### 3. Настройка окружения

Необходимо задать переменные окружения:

```bash
INSTANCE_ID= #idInstance
API_TOKEN= #apiTokenInstance
API_URL= #apiUrl
CHAT_ID= #chatId
```

Пример (Linux/macOS):

```bash
export INSTANCE_ID=xxx
export API_TOKEN=xxx
export API_URL=xxx
export CHAT_ID=xxx
```

Или создать .env файл в соответствии c примером: .env.example

---

### 4. Запуск тестов

Через Makefile:

```bash
make test
```

Без кеша:

```bash
make test-no-cache
```

Или напрямую:

```bash
go test -v ./tests
```

---

## Что проверяют тесты

* успешные (200) и ошибочные (400) ответы
* обязательные параметры
* авторизацию инстанса
* обработку ошибок API
* rate limiting (429)

---

## Примечания

* Тесты являются интеграционными и зависят от внешнего API
* Для обработки 429 используется повтор запросов (retry)
* Перед запуском тестов инстанс должен быть авторизован
