# Go API Dockerized Using Gin

This is a simple Go API implemented using the Gin framework and containerized using Docker.

## Prerequisites

- Docker: Ensure that you have Docker installed on your system.

## Getting Started

To get started with the Go API Dockerized, follow the instructions below.

### Clone the Repository

```bash
git clone https://github.com/your-username/your-repo.git
cd your-repo
```

Build the Docker Image
```bash
docker build -t go-api-dockerized .
```

Run the Docker Container
```bash
docker run -p 8080:8080 go-api-dockerized
```

The API will be accessible at http://localhost:8080.

### API Endpoints
The following API endpoints are available:
GET /albums: Get a list of all albums.
POST /albums: Add a new album.
GET /albums/{id}: Get details of a specific album.


### Configuration
The Go API Dockerized can be configured using environment variables:
MYAPI_PASSWORD: The password required to access the API (default: "pass").
MYAPI_HOST: The host IP or domain to bind the API server to (default: "0.0.0.0").
MYAPI_PORT: The port number to listen on (default: "8080").

## Contributing
Contributions to the Go API Dockerized are welcome! If you find a bug or want to add a new feature, please open an issue or submit a pull request.

