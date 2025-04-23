# Guide

## How to run Dev

- Create .env file with this config in root directory.

```bash
# Echo app
APP_ENV=development
TIMEZONE=UTC
LOCAL_TIMEZONE=Asia/Bangkok

# Postgres
DATABASE_HOST=postgres-cart
DATABASE_USER=postgres
DATABASE_PASSWORD=postgres
DATABASE_NAME=cart_db
DATABASE_HOST_PORT=5434
DATABASE_DOCKER_PORT=5432

# Docker
COMPOSE_PROJECT_NAME=demo-cart-service
APP_BUILD_CONTEXT=../../

# NATS
NATS_URL=nats://nats:4222 # Docker service name:port
NATS_TOKEN=#platong1234

# gRPC
GRPC_SHOP_PRODUCT_ENDPOINT="demo-shop-product-service-app-1:50051"
GRPC_PAYMENT_ENDPOINT="demo-payment-service-app-1:50051"
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

## NATS JetStream

- <https://github.com/nats-io/nats.go/blob/main/jetstream/README.md>
