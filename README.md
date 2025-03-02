# Evermos Internship Project

## Overview

This project is an e-commerce API built using Golang and MySQL, following the Clean Architecture pattern. It includes JWT authentication, role-based access control, pagination, and filtering.

## Technologies Used

- **Golang** (Gin framework)
- **MySQL** (GORM ORM)
- **JWT Authentication**
- **Postman** for API testing

## Features

- **User Management** (Register, Login, Role-based access control)
- **Store Management** (CRUD operations, auto-created on user registration)
- **Address Management** (User can manage multiple addresses)
- **Category Management** (Admin-only feature)
- **Product Management** (CRUD operations, image upload)
- **Transaction System** (Purchase handling with product logs)
- **Pagination & Filtering** (Implemented in multiple endpoints)

## API Documentation

### **Authentication**

#### 1. Register User

`POST /register`

```json
{
  "name": "John Doe",
  "email": "johndoe@example.com",
  "password": "password123"
}
```

**Response:**

```json
{
  "message": "User registered successfully",
  "user_id": 1
}
```

#### 2. Login User

`POST /login`

```json
{
  "email": "johndoe@example.com",
  "password": "password123"
}
```

**Response:**

```json
{
  "token": "your_jwt_token"
}
```

### **Product APIs**

#### 3. Create Product (Authenticated & Authorized: Admin)

`POST /products`

```json
{
  "name": "Laptop XYZ",
  "price": 1200.50,
  "category_id": 1,
  "stock": 10,
  "description": "High-end gaming laptop"
}
```

**Response:**

```json
{
  "message": "Product created successfully",
  "product_id": 1
}
```

#### 4. Get All Products (With Pagination & Filtering)

`GET /products?page=1&limit=10&category=Electronics`
**Response:**

```json
{
  "products": [
    {
      "id": 1,
      "name": "Laptop XYZ",
      "price": 1200.50,
      "category": "Electronics"
    }
  ],
  "total_pages": 5
}
```

### **Transaction APIs**

#### 5. Create Transaction (Authenticated User)

`POST /transactions`

```json
{
  "shipping_address": "123 Main St, City",
  "total_price": 150.75,
  "invoice_code": "INV123456",
  "payment_method": "Credit Card",
  "status": "Pending",
  "total_amount": 3
}
```

**Response:**

```json
{
  "message": "Transaction created successfully",
  "transaction_id": 1
}
```

#### 6. Get All Transactions (Admin Only)

`GET /transactions`
**Response:**

```json
{
  "transactions": [
    {
      "id": 1,
      "user_id": 6,
      "total_price": 150.75,
      "status": "Pending"
    }
  ]
}
```

#### 7. Get User Transactions (Authenticated User)

`GET /transactions/my`
**Response:**

```json
{
  "transactions": [
    {
      "id": 1,
      "total_price": 150.75,
      "status": "Pending"
    }
  ]
}
```

### **Pagination & Filtering**

Most `GET` endpoints support pagination:

- `?page=1&limit=10` → Fetch 10 results per page
- `?status=pending` → Filter transactions by status

## Setup Instructions

1. Clone the repository

```sh
git clone https://github.com/yourusername/yourrepo.git
cd yourrepo
```

2. Install dependencies

```sh
go mod tidy
```

3. Configure `.env` file with your database credentials
4. Run the server

```sh
go run main.go
```

## Conclusion

This project demonstrates a full-fledged e-commerce API with authentication, role-based access, and a transaction system. Future improvements may include better logging, an admin dashboard, and additional security features.

---

**Author:** Mitchel M. Affandi\
**Internship at:** Evermos

