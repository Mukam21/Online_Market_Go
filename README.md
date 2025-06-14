# Онлайн Маркет на Go с Kafka

Проект — простой онлайн-маркет на Golang с REST API на Gin и системой обработки заказов через Kafka.

---

## Основные возможности

- CRUD операции с товарами (создание, чтение, обновление, удаление)

- Поиск товаров

- Регистрация и аутентификация пользователей

- Создание заказов с асинхронной обработкой через Kafka

- Просмотр профиля пользователя

---

## Архитектура

- **Gin** — HTTP фреймворк для API

- **Kafka** — брокер сообщений для асинхронной обработки заказов

- **PostgreSQL** — основная база данных (подключение через `pkg/database`)

- **Микросервисный подход:** HTTP-сервер и Kafka consumer работают отдельно

---

## Быстрый старт

### 1. Запуск Kafka и Zookeeper через Docker

```bash
docker-compose up -d

2. Настройка базы данных
-  Подключитесь к PostgreSQL (конфигурация в pkg/database)

-  Выполните миграции (если есть)

3. Запуск сервера

  go run cmd/main.go

API эндпоинты

GET	/api/products	Получить список товаров

GET	/api/products/:id	Получить товар по ID

POST	/api/products	Создать новый товар

PUT	/api/products/:id	Обновить товар по ID

DELETE	/api/products/:id	Удалить товар по ID

POST	/api/products/:id/purchase	Купить товар

GET	/api/products/search	Поиск товаров

POST	/api/register	Регистрация пользователя

POST	/api/login	Вход пользователя

GET	/api/user_profile	Профиль пользователя

POST	/api/orders	Создать заказ (через Kafka)

Как работает Kafka в проекте

-  Заказы публикуются в Kafka топик orders

-  Отдельный consumer читает сообщения и логирует детали заказа

-  Позволяет масштабировать обработку заказов асинхронно и эффективно

Требования

-  Go 1.20+

-  Docker и docker-compose

-  Kafka и Zookeeper (можно поднять через docker-compose)

-  PostgreSQL
