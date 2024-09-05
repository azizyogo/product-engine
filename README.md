# product-engine

## About

This service provides functionality related to Products for a marketplace system. There are many potential features that can be developed in this service, such as:

- Product Display
- Product Reviews
- Product Discounts / Vouchers
- Product Management
- etc.

However, in the current version, only Product Management has been developed. Product Management is a feature that allows users or stores to manage their products, including inserting new products, viewing details, editing, and deleting existing products.

## Requirement

- Go 1.22.4 or higher
- Database MongoDB
- Docker

## Installation

1. Clone Repository

```bash
git clone https://github.com/azizyogo/product-engine.git
```

2. Copy env in project root and setting your environment

```bash
cp .env.example .env.yaml
```

3. Run with Docker

```bash
make docker-run
```

4. Run Migration *excute this command twice to run migrations file in directory ./migrations

```bash
make migrate-up
```

## CURL

- Login
```
curl --location 'localhost:8080/login' \
--header 'Content-Type: application/json' \
--data '{
    "username": "samsung.official",
    "password": "pwd000"
}'
```

- Get Product Detail
```
curl --location --request GET 'localhost:8080/product' \
--header 'Authorization: Bearer <put token here>' \
--header 'Content-Type: application/json' \
--data '{
    "product_id": "You can find the product ID in MongoDB after running the migration"
}'
```

- Insert New Product
```
curl --location 'localhost:8080/product' \
--header 'Authorization: Bearer <put token here>' \
--header 'Content-Type: application/json' \
--data '{
    "user_id": "Select the user ID to which you want to add a product. Fnd the user ID in MongoDB after running the migration",
    "name": "ZD Wolverine",
    "category": "",
	"condition": 1,
	"desc": "",
	"price": 400000,
	"status": 1,
	"stock": 10,
	"specifications": [
        {"Age": "15"},
        {"Brand": "ZD Toys"}
    ]
}'
```

- Delete Product
```
curl --location --request DELETE 'localhost:8080/product/{Product ID}' \
--header 'Authorization: Bearer <put token here>'
```

- Update Cake
```
curl --location --request PUT 'localhost:8080/product' \
--header 'Authorization: Bearer <put token here>' \
--header 'Content-Type: application/json' \
--data '{
        "id": "Put the Product ID which you want to edit",
        "user_id": "Put the User ID which you want to edit",
        "name": "Samsung Z Flip 6",
        "category": "",
        "condition": 1,
        "desc": "the newest flip phone from samsung",
        "price": 15000000,
        "status": 1,
        "stock": 100,
        "specifications": [
            {
                "Brand": "Samsung"
            },
            {
                "Warranty": "SEIN"
            },
            {
                "RAM": "6GB"
            }
        ]
    }'
```

## Testing Guides

### Product Management
Since this featues meant to user / store to manage their product, you have to login first.

1. Login to get access token
2. Put token in Authorization header of each Product Management endpoints

## Notes

if you encounter this kind of error when doing migration up
```
error: database driver: unknown driver mongodb (forgotten import?)
```
run this command
```
go install -tags "mongodb" github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```