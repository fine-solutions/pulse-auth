[![Product Deploy](https://github.com/fine-solutions/pulse-back/actions/workflows/product-deploy.yaml/badge.svg)](https://github.com/fine-solutions/pulse-back/actions/workflows/product-deploy.yaml)

# Микросервисы (бэкенд) для программного комплекса Пульс

## Установка и запуск

1. [Установить](https://docs.docker.com/engine/install/), если она ещё не установлена, службу Docker (Docker Compose). 
2. Скопировать файл `.env.example` в `.env`:

```bash
cp .env.example .env
```

3. Установить значения переменных окружения, например:

Список переменных окружения:

- `DB_USER` — основной пользователь БД с правами администратора;
- `DB_PASSWORD` — пароль основного пользователя;
- `DB_NAME` — наименование базы данных;
- `DB_DATA` — путь к файлу с данными, которые хранятся в томе Docker;
- `DB_PORT` — порт базы данных;
- `JWT_TOKEN_SALT` — «соль» (добавка к уникальному хешу) токена авторизации.

```bash
# Настройка доступа к базе данных
DB_USER=pulse
DB_PASSWORD=pulse
DB_NAME=pulse
DB_DATA=/var/lib/postgresql/data/pgdata
DB_PORT=5432

# Настройка авторизации
JWT_TOKEN_SALT=pulse
```

4. Запустить микросервисы, предварительно собрав их docker-образы:

```bash
docker compose down
docker compose build
docker compose --env-file .env up --detach
```

## Разработка

### Отладка совместной работы микросервисов

1. Выполнить пункты раздела **Установка и запуск**.
1. Для проверки совместной работы микросервисов после обновления кода, необходимо повторить пункт 4 раздела **Установка и запуск**.

### Отладка работы микросервисов по отдельности

1. [Запуск и отладка](auth/README.md) микросервиса для авторизации.
