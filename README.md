🛒 Online Market API (на Golang)
Простое RESTful API для онлайн-маркета, реализованное с помощью:

⚙️ Go (Golang)

🌐 Gin (веб-фреймворк)

🗄️ GORM (ORM для работы с PostgreSQL)

🔐 JWT авторизация

🐘 PostgreSQL

📦 Установка и запуск
1. Клонируй репозиторий
git clone https://github.com/Mukam21/Online_Market_Go.git

2. Настрой .env файл
Создай файл .env

3. Запусти проект
go run cmd/main.go

🚀 API Эндпоинты
🔐 Аутентификация
POST /api/register — регистрация пользователя

POST /api/login — логин и получение JWT токена

🛒 Товары
Все маршруты /api/products защищены JWT (кроме GET)

GET /api/products — получить все товары

GET /api/products/:id — получить товар по ID

POST /api/products — создать товар (требуется токен)

PUT /api/products/:id — обновить товар (требуется токен)

DELETE /api/products/:id — удалить товар (требуется токен)

👤 Пользовательский профиль
GET /api/user_profile
Описание: Получить данные текущего авторизованного пользователя по JWT-токену.

Требуется JWT-токен: ✅



🧠 Автор
Mukam21
GitHub: @Mukam21

