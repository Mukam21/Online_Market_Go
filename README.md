Онлайн-маркет, реализованный на Go с использованием Gin, PostgreSQL и JWT для аутентификации.

## Возможности

- Регистрация и авторизация пользователей (JWT)

- CRUD для товаров

- Поиск товаров по имени

- Покупка товара

- Просмотр профиля пользователя

## Технологии

- Go (Golang)

- Gin (HTTP фреймворк)

- PostgreSQL (СУБД)

- Gorm (ORM)

- JWT (JSON Web Tokens)

## Установка и запуск

1. Клонируйте репозиторий:

   ```bash
   git clone https://github.com/Mukam21/Online_Market_Go.git

2. Настройте базу данных PostgreSQL:

-  Создайте базу данных и пользователя.

-  Обновите настройки подключения в config (например, .env или в коде).

3. Запустите приложение:

   go run cmd/main.go

4. Сервер запустится на http://localhost:8080

API
Пользователи:
-  POST /api/register — регистрация

-  POST /api/login — вход (возвращает JWT токен)

-  GET /api/user_profile — получение профиля (требует JWT)

Товары:
-  GET /api/products — список всех товаров

-  GET /api/products/:id — получить товар по ID

-  POST /api/products — создать товар

-  PUT /api/products/:id — обновить товар

-  DELETE /api/products/:id — удалить товар

-  GET /api/products/search?q=название — поиск товаров по имени

-  POST /api/products/:id/purchase — покупка товара


🧠 Автор
Mukam21
GitHub: @Mukam21

