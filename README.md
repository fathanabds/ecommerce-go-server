# E-Commerce Go Server

This project is an e-commerce backend API built with Go and the Gin framework.

## Authors

- [@fathanabds](https://github.com/fathanabds)

## Environment Variables

To run this project, you will need to add the following environment variables to your .env file

`JWT_SECRET`

## Features

- User Authentication: Register a New User, Login a User and Generate a JWT Token
- Product Management: Get a List of All Products, Add a New Product (Admin Only)
- Cart Management: Add a Product to the User’s Cart, Get the User’s Cart

## Run Locally

Clone the project

```bash
  git clone https://github.com/fathanabds/ecommerce-go-server.git
```

Go to the project directory

```bash
  cd ecommerce-go-server
```

Install dependencies

```bash
  go mod tidy
```

Start the server

```bash
  go run main.go
```
