# Go Clean Architecture

This project is an example implementation of REST API using clean architecture in Go (Golang).

### Built With

- [Echo](https://echo.labstack.com/)
- [GORM](https://gorm.io/)

## Getting Started

To get a local copy up and running follow these simple example steps.

### Prerequisites

To run the project you need:

- Go (Golang) (1.16.4 has been confirmed to work)
- Docker compose

## Quick Start

Run the project locally according to the following command.

1. Start the containers in the background

   ```sh
   docker-compose up -d
   ```

2. Change directory to service

   ```sh
   cd service
   ```

3. Generate RSA key pair (key pair is used to generate and validate JWT Token)

   ```sh
   openssl genrsa -out app.rsa && openssl rsa -in app.rsa -pubout > app.rsa.pub
   ```

4. Compile and run Go program
   ```sh
   go run main.go
   ```

<!-- LICENSE -->

## License

Distributed under the MIT License. See `LICENSE` for more information.
