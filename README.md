# Evermos Internship Project

This project is a **Golang-based e-commerce API** using **MySQL**, built with **clean architecture**. It includes JWT authentication, pagination, filtering, and user access restrictions.

## Features
- User authentication with JWT
- Product and category management
- Transaction processing
- Role-based access control
- Pagination and filtering

## Tech Stack
- **Golang**
- **MySQL**
- **GORM**
- **Gin (for API routing)**

## Installation

### Prerequisites
- Go installed (`>=1.18`)
- MySQL installed and running

### Setup Steps
1. Clone the repository:
   ```sh
   git clone https://github.com/your-repo/EvermosInternship.git
   cd EvermosInternship
   ```
2. Create a `.env` file and configure your database:
   ```sh
   DB_USER=root
   DB_PASSWORD=yourpassword
   DB_HOST=localhost
   DB_PORT=3306
   DB_NAME=go_ecommerce
   JWT_SECRET=your_secret_key
   ```
3. Install dependencies:
   ```sh
   go mod tidy
   ```
4. Run database migrations:
   ```sh
   go run main.go migrate
   ```
5. Start the server:
   ```sh
   go run main.go
   ```

## API Endpoints

### Authentication
- **POST /register** – Register a new user
- **POST /login** – Login and get JWT token

### Products
- **GET /products** – Get list of products
- **POST /products** – Create a new product (Admin only)

### Transactions
- **POST /transactions** – Create a new transaction
- **GET /transactions** – Get user transactions
- **GET /transactions/:id** – Get transaction details

## Testing with Postman
- Import the provided `postman_collection.json`
- Set the **Authorization Bearer Token** before making API calls

## Contributors
- **Your Name** – Backend Developer

## License
This project is licensed under the **MIT License**.

