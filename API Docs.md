# API Documentation

## User Endpoints

### Register User

- **URL:** `/register`
- **Method:** `POST`
- **Description:** Register a new user.
- **Request Body:**
  ```json
  {
    "username": "string",
    "email": "string",
    "password": "string",
    "role": "string"
  }
  ```
- **Response:**
  - **201 Created:**
    ```json
    {
      "message": "User registered successfully"
    }
    ```
  - **400 Bad Request:**
    ```json
    {
      "error": "Validation error message"
    }
    ```

### Login User

- **URL:** `/login`
- **Method:** `POST`
- **Description:** Login a user.
- **Request Body:**
  ```json
  {
    "email": "string",
    "password": "string"
  }
  ```
- **Response:**
  - **200 OK:**
    ```json
    {
      "message": "Login successful",
      "token": "jwt_token"
    }
    ```
  - **400 Bad Request:**
    ```json
    {
      "error": "Validation error message"
    }
    ```
  - **401 Unauthorized:**
    ```json
    {
      "error": "Invalid email or password"
    }
    ```

## Product Endpoints

### Get Products

- **URL:** `/products`
- **Method:** `GET`
- **Description:** Get all products.
- **Headers:**
  - `Authorization: Bearer <jwt_token>`
- **Response:**
  - **200 OK:**
    ```json
    {
      "message": "Products fetched successfully",
      "products": [
        {
          "id": 1,
          "name": "Product A",
          "price": 10000
        },
        {
          "id": 2,
          "name": "Product B",
          "price": 20000
        }
      ]
    }
    ```
  - **200 OK (No products available):**
    ```json
    {
      "message": "No products available",
      "products": []
    }
    ```

### Add Product

- **URL:** `/products`
- **Method:** `POST`
- **Description:** Add a new product (Admin only).
- **Headers:**
  - `Authorization: Bearer <jwt_token>`
- **Request Body:**
  ```json
  {
    "name": "string",
    "price": "number"
  }
  ```
- **Response:**
  - **201 Created:**
    ```json
    {
      "message": "Product added successfully",
      "product": {
        "id": 3,
        "name": "Product C",
        "price": 30000
      }
    }
    ```
  - **400 Bad Request:**
    ```json
    {
      "error": "Validation error message"
    }
    ```
  - **403 Forbidden:**
    ```json
    {
      "error": "You do not have permission"
    }
    ```

## Cart Endpoints

### Add to Cart

- **URL:** `/carts`
- **Method:** `POST`
- **Description:** Add a product to the cart.
- **Headers:**
  - `Authorization: Bearer <jwt_token>`
- **Request Body:**
  ```json
  {
    "product_id": "number",
    "quantity": "number"
  }
  ```
- **Response:**
  - **200 OK:**
    ```json
    {
      "message": "Product added to cart",
      "cart": [
        {
          "product_id": 1,
          "quantity": 2
        }
      ]
    }
    ```
  - **400 Bad Request:**
    ```json
    {
      "error": "Invalid input"
    }
    ```
  - **404 Not Found:**
    ```json
    {
      "error": "Product not found"
    }
    ```

### Get Cart

- **URL:** `/carts`
- **Method:** `GET`
- **Description:** Get the user's cart.
- **Headers:**
  - `Authorization: Bearer <jwt_token>`
- **Response:**
  - **200 OK:**
    ```json
    {
      "message": "Cart fetched successfully",
      "cart": [
        {
          "product_id": 1,
          "quantity": 2
        }
      ]
    }
    ```
  - **200 OK (Cart is empty):**
    ```json
    {
      "message": "Cart is empty",
      "cart": []
    }
    ```

## Middleware

### JWT Authentication Middleware

- **Description:** Validates the JWT token and sets user information in the context.
- **Usage:** Applied to routes that require authentication.

### Role Authorization Middleware

- **Description:** Verifies the user's role.
- **Usage:** Applied to routes that require specific roles (e.g., admin).
