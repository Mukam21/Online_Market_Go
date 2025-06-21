 Онлайн-Маркет с Kafka, Gin, GORM и PostgreSQL
 
  Описание:

 Это простой пример онлайн-маркета, реализованный на Go с использованием:

     - Gin — веб-фреймворк для создания REST API

     - GORM — ORM для работы с базой данных (PostgreSQL)

     - Kafka — брокер сообщений для асинхронной обработки заказов

     - PostgreSQL — база данных

Заказы создаются через REST API и отправляются в Kafka. Консьюмер слушает Kafka, сохраняет заказы в базу и обновляет количество товаров.

  Функциональность:
  
- CRUD для товаров

- Регистрация и авторизация пользователей

- Поиск товаров

- Создание заказов с помощью Kafka (продюсер)

- Обработка заказов (консьюмер Kafka) с сохранением в базу и уменьшением количества товара

  Требования:
  
- Go 1.20+

- Docker и Docker Compose (для Kafka и Zookeeper)

- PostgreSQL (локально или через Docker)

  Запуск проекта:
  
1. Запуск Kafka и Zookeeper через Docker Compose

            docker-compose up -d

2. Настройка базы данных:

- Создай базу PostgreSQL

- В .env файле (или в коде) укажи настройки подключения к БД

3. Миграции:

- Миграции таблиц запускаются автоматически при старте приложения.

4. Запуск сервера (продюсера и API):

            go run cmd/main.go

5. Запуск консьюмера Kafka (обработчик заказов)

  В отдельном терминале:

            go run cmd/consumer/main.go

  Примеры запросов:

Регистрация пользователя:

     **** POST /api/register

          Content-Type: application/json

          {
      
            "username": "user1",
        
           "password": "password123"
        
          } ****

Авторизация:

    **** POST /api/login
 
         Content-Type: application/json

        {
      
           "username": "user1",
  
           "password": "password123"

        } ****

В ответ получишь JWT токен, который нужно передавать в заголовке Authorization для защищенных эндпоинтов:

  Authorization: Bearer <token>

Получить список товаров:

       GET /api/products

Создать заказ (отправляется в Kafka):

Например :

      **** POST /api/orders
      
            Content-Type: application/json
      
           {
      
               "user_id": 1,
   
               "product_id": 5,
  
               "quantity": 2

           } ****

Структура проекта:

     - cmd/main.go — точка входа для API + Kafka продюсера

     - cmd/consumer/main.go — Kafka консьюмер, который обрабатывает заказы

     - pkg/database/ — подключение к базе и миграции 

     - pkg/models/ — модели БД (User, Product, Order)

     - pkg/handlers/ — HTTP хендлеры для API

     - pkg/routes/ — маршруты API

     - pkg/middleware/ — JWT авторизация

    - pkg/jwt/ — JWT авторизация

    - docker-compose.yml — конфигурация для Kafka и Zookeeper
