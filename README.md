# Bomb Party

## Features

- **JWT Authentication** : Secure login system using JWT tokens  
- **User Management** : Full CRUD for users accounts  
- **Team Management** : Full CRUD for teams  
- **Game Management** : Full CRUD for games  
- **Bomb Management** : Full CRUD for bombs  
- **Inventory Management** : Full CRUD for inventories  
- **Swagger Documentation** : Interactive interface to test the API  
- **SQLite Database** : Persistent storage using GORM  

## Configuration

The application uses SQLite as the default database.  
The file `bomb-party.db` is automatically created on first launch.

The configuration (in the `config` package) initializes:

- Database connection  
- Repositories for each entity  
- Automatic schema migrations  

## Technologies

- **Go**  
- **Chi Router** – Lightweight and fast HTTP router  
- **GORM** – ORM for Go  
- **SQLite** – Embedded database  
- **JWT (golang-jwt)** – Secure authentication  
- **bcrypt** – Password hashing  
- **Swagger** – Automatic API documentation  
- **http-swagger** – Swagger UI  
- **Bruno** – API route testing  

## Installation

Requires **Go 1.24+**

### Steps

1. **Clone the repository**

```bash
git clone https://github.com/Y-BombParty/BombParty-Api.git
cd BombParty-Api
```

2. **Install dependencies**

```bash
go get
swag init
```

3. **Run the application**

```bash
go run main.go
```

API available at:  
`http://localhost:7774`

### Swagger Documentation

Once the server is running:

```
http://localhost:8080/swagger/index.html
```
## Routes

POST   /api/v1/users/login
POST   /api/v1/users/register
GET    /api/v1/users/
GET    /api/v1/users/{id}
PUT    /api/v1/users/{id}
DELETE /api/v1/users/{id}

GET    /api/v1/bombs/
POST   /api/v1/bombs/
GET    /api/v1/bombs/user/{userId}
GET    /api/v1/bombs/{id}
PUT    /api/v1/bombs/{id}
DELETE /api/v1/bombs/{id}

POST   /api/v1/inventory/add
GET    /api/v1/inventory/init
GET    /api/v1/inventory/inventory

GET    /api/v1/teams/
POST   /api/v1/teams/
PUT    /api/v1/teams/{id}
DELETE /api/v1/teams/{id}
GET    /api/v1/teams/{id}
 
## API Documentation

The API uses Swagger/OpenAPI.  
Each endpoint includes:

- Feature description  
- Required parameters  
- Request/response formats  
- HTTP status codes  

## Members

**Emmanuel Yohore**, **Aurélien DUGAS**, **Jean-Baptiste BODUSSEAU**, **Mathis SILOTIA**
