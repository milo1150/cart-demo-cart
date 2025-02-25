# Guide

## How to run Dev

- Create .env file with this config in root directory.

```bash
APP_ENV=development
DATABASE_HOST=postgres-shop-product
DATABASE_USER=postgres
DATABASE_PASSWORD=postgres
DATABASE_NAME=cart_db
DATABASE_HOST_PORT=5433
DATABASE_DOCKER_PORT=5432
TIMEZONE=UTC
LOCAL_TIMEZONE=Asis/Shanghai

COMPOSE_PROJECT_NAME=demo-cart-service
```

- For first time.

```bash
cd scripts && chmod +x dev-start.sh && ./dev-start.sh
```

- Later

```bash
cd scripts && ./dev-start.sh
```

## Database CLI

```bash
pgcli postgres://postgres:postgres@127.0.0.1:5434/cart_db
```
