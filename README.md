<h1 align="center">RESTAURANT MANAGEMENT BACKEND</h1>

<div align="center">
This backend service is designed for a restaurant management system built with Golang. It provides APIs to manage various aspects of the restaurant, including menu items, orders, invoices, table management, and user authentication.
</div>

## Table of Contents

- [Features](#features)
- [Installation](#installation)
- [Configuration](#configuration)
- [Usage](#usage)
- [Database Management](#database-management)
- [API Endpoints](#api-endpoints)
- [License](#license)

## Features

- JWT-based authentication and authorization
- RESTful API with Gin framework
- Input validation using go-playground/validator
- MongoDB for database operations

## Installation

### Prerequisites

Before you begin, ensure you have the following installed on your machine:

- [Golang](https://golang.org/dl/) 1.20 or later
- [MongoDB](https://www.mongodb.com/try/download/community) 4.4 or later
- [Docker](https://docs.docker.com/engine/install/)
- [Git](https://git-scm.com/)

### Clone the Repository

```bash
git clone https://github.com/Nhathuy1305/AdminGo.git
cd AdminGo
```

### Install Dependencies

This project uses Go modules to manage dependencies. You can install the required dependencies using the following command:

```bash
go mod tidy
```

### Environment Variables

Create a `.env` file in the root directory of your project and configure the following environment variables:

```env
PORT=8000 #Optional
MONGO_INITDB_ROOT_USERNAME=
MONGO_INITDB_ROOT_PASSWORD=
MONGO_HOST=
MONGO_PORT=
```

### Run the Application

Create the Docker Images with these commands:
```bash
docker run -d --name mongo-demo \
-e MONGO_INITDB_ROOT_USERNAME=your-username \
-e MONGO_INITDB_ROOT_PASSWORD=your-password \
-p 27018:27017 \
mongo:6.0
```
* Remember to change `mongo-demo`, `your-username` and `your-password` with your image, root username and password of your database.

Run the application by the following command:
```bash
go run main.go
```

The backend service will start on `http://localhost:8000`.

## Usage

### Building the Application

To build the application, run the following command:

```bash
go build -o admingo
```

### Running the Built Application

```bash
./admingo
```

## Database Management

1. Access the MongoDB Image using this command:
```bash
docker exec -it mongo-demo /bin/bash
```
2. Access the MongoDB shell:
```bash
mongosh -u your-username -p your-password
```
3. List all databases:
```bash
show dbs
```
4. Access a specific database:
```bash
use your_database_name
```
5. List all collections (tables) in the database:
```bash
show collections
```
6. Show data in a specific collection:
```bash
db.your_collection_name.find().pretty()
```

For example:
```bash
docker exec -it mongo-demo /bin/bash

mongosh -u mongotest -p 123456

show dbs

use restaurant

show collections

db.user.find().pretty()
```

## API Endpoints

### Users
- `GET /users` - Retrieve a list of users.
- `GET /users/:user_id` - Retrieve a specific user by ID.
- `POST /users/signup` - Register a new user.
- `POST /users/login` - Authenticate a user and obtain a JWT token.

### Foods
- `GET /foods` - Retrieve a list of food items.
- `GET /foods/:food_id` - Retrieve a specific food item by ID.
- `POST /foods` - Add a new food item.
- `PATCH /foods/:food_id` - Update an existing food item.

### Menus
- `GET /menus` - Retrieve a list of menus.
- `GET /menus/:menu_id` - Retrieve a specific menu by ID.
- `POST /menus` - Add a new menu.
- `PATCH /menus/:menu_id` - Update an existing menu.

### Tables
- `GET /tables` - Retrieve a list of tables.
- `GET /tables/:table_id` - Retrieve a specific table by ID.
- `POST /tables` - Add a new table.
- `PATCH /tables/:table_id` - Update an existing table.

### Orders
- `GET /orders` - Retrieve a list of orders.
- `GET /orders/:order_id` - Retrieve a specific order by ID.
- `POST /orders` - Place a new order.
- `PATCH /orders/:order_id` - Update an existing order.

### Order Items
- `GET /orderItems` - Retrieve a list of order items.
- `GET /orderItems/:orderItem_id` - Retrieve a specific order item by ID.
- `GET /orderItems-order/:order_id` - Retrieve order items by order ID.
- `POST /orderItems` - Add a new order item.
- `PATCH /orderItems/:orderItem_id` - Update an existing order item.

### Invoices
- `GET /invoices` - Retrieve a list of invoices.
- `GET /invoices/:invoice_id` - Retrieve a specific invoice by ID.
- `POST /invoices` - Create a new invoice.
- `PATCH /invoices/:invoice_id` - Update an existing invoice.

## Dependencies

This project uses the following Go packages:

- [github.com/dgrijalva/jwt-go v3.2.0+incompatible](https://pkg.go.dev/github.com/dgrijalva/jwt-go)
- [github.com/gin-gonic/gin v1.10.0](https://pkg.go.dev/github.com/gin-gonic/gin)
- [github.com/go-playground/validator/v10 v10.20.0](https://pkg.go.dev/github.com/go-playground/validator/v10)
- [go.mongodb.org/mongo-driver v1.16.0](https://pkg.go.dev/go.mongodb.org/mongo-driver)
- [golang.org/x/crypto v0.23.0](https://pkg.go.dev/golang.org/x/crypto)

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.