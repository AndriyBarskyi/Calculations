# Calculations Service

This is a simple gRPC service written in Go that performs basic arithmetic operations such as addition and division.

## Prerequisites

- Go 1.16 or higher
- Docker

## Getting Started

Clone the repository to your local machine:

```bash
git clone https://github.com/AndriyBarskyi/Calculations.git
```

Navigate to the project directory:

```bash
cd Calculations
```

## Running the Application Locally

To run the application locally, execute:

```bash
go run cmd/main.go
```

The service will start and listen on port 8089.

## Running the Application in Docker

To build the Docker image, execute:

```bash
docker build -t calculations .
```

To run the application in a Docker container, execute:

```bash
docker run -p 8089:8089 calculations
```

The service will start and listen on port 8089.

## Testing

To run the tests, execute:

```bash
go test ./...
```

## API

The service provides the following gRPC methods:

- `Add`: Adds two numbers.
- `Divide`: Divides the first number by the second number.

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License

[MIT](https://choosealicense.com/licenses/mit/)
