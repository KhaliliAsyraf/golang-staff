# 🧑‍💼 Employee Leaves System

A simple **Employee Leave Management System** built with **Golang (v1.23.2)** that allows employees to apply for leaves.  
This project includes authentication, department management, and employee management features — running seamlessly with **Docker**.

---

## 🧱 Tech Stack

- **Backend:** Golang 1.23.2  
- **Containerization:** Docker  
- **Database:** (Configured within Docker)  
- **API Testing:** Postman  

> 💡 You can either use Docker or run Go directly on your machine — just ensure Golang is installed and configured properly.

---

## 🚀 Getting Started

### 1. Prerequisites

Make sure you have the following installed:

- [Docker](https://www.docker.com/)  
- [Docker Compose](https://docs.docker.com/compose/)  
- [Postman](https://www.postman.com/downloads/) (for API testing)

Alternatively, if not using Docker, ensure:
- [Golang 1.23.2](https://go.dev/dl/) is installed
- Database credentials are properly configured in your `.env` file

---

### 2. Start the Application

If you prefer to run the project locally instead of using Docker, follow these steps:

1. Run the setup script to handle migrations and seed data:
```
./setup.sh
```

2. Once setup is complete, start the Go application:
```
go run .
```

Else if running application with Docker:

Open your terminal and run:

```bash
docker compose up -d --build
```

This command will:

- Build and start the Docker container

- Set up the Golang environment

- Run database migrations

- Seed initial data (including admin credentials)


### 3. Verify Database Connection

Once the containers are up, open any DBMS tool (e.g., DBeaver, TablePlus, or VSCode SQL extension) to confirm:

- The database has been created successfully

- Tables and seed data exist

---

### 4. 🔐 API Usage (via Postman)

### Base URL

```bash
http://localhost:8080
```

### 1️⃣ Login

Endpoint: POST /login

Description: Authenticate as admin to obtain a token.

Request Body:
```json
{
  "email": "admin@example.com",
  "password": "password"
}
```

Response example:
```json
{
  "token": "your-jwt-token"
}
```

👉 Copy the token for subsequent requests.

### 2️⃣ Get Departments

Endpoint: GET /departments

Description: Retrieve a list of available departments.

Headers:
```bash
Authorization: Bearer <your-jwt-token>
```

Response Example:
```json
[
  {
    "id": 1,
    "name": "Human Resources"
  },
  {
    "id": 2,
    "name": "Engineering"
  }
]
```