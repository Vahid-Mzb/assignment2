# Key-Value Store Database in Go

A simple persistent key-value store with HTTP REST API endpoints, implemented in Go.

##  Table of Contents
- [Features](#features)
- [Prerequisites](#prerequisites)
- [Installation & Setup](#installation--setup)
  - [Option 1: Run Locally](#option-1-run-locally)
  - [Option 2: Run with Docker](#option-2-run-with-docker)
- [Error Handling](#error-handling)
- [Usage Examples](#usage-examples)

## Features

- **PUT /objects** - Store key-value pairs
- **GET /objects/{key}** - Retrieve values by key
- **Persistent storage** using JSON file
- **Dockerized** for easy using
- Data persists across application restarts

## Prerequisites

- Go 1.25.3 or higher (for local development)
- Docker and Docker Compose (for containerized deployment)

## Installation & Setup

### Option 1: Run Locally


1. Initialize Go modules:
```bash
go mod download
```

2. Run the application:
```bash
go run main.go
```

The server will start on `http://localhost:8080`

### Option 2: Run with Docker
---
### Installing Docker on Linux

To **download and install Docker**, run the following command in your terminal:

```bash
curl -fsSL get.docker.com -o get-docker.sh && sh get-docker.sh
```

---
### Using a VPN or Mirror for Docker Images

To **download Docker images**, you either need a **VPN** or you can use a **mirror**, such as the **AbrArvan mirror**.

To configure Docker to use the AbrArvan mirror, run this command in your terminal:

```bash
sudo bash -c 'cat > /etc/docker/daemon.json <<EOF
{
  "insecure-registries": ["https://docker.arvancloud.ir"],
  "registry-mirrors": ["https://docker.arvancloud.ir"]
}
EOF'
```

---

### Applying the Changes

After updating the configuration, run the following commands to apply the changes:

```bash
docker logout
sudo systemctl restart docker
```

---

 Docker should now be configured and ready to pull images using the AbrArvan mirror.
 
 ### Now let's run the code using Docker:
 
1. Build and run using Docker Compose:
```bash
docker compose up --build
```

2. The server will be available at `http://localhost:8080`

3. To run in detached mode:
```bash
docker compose up -d
```

4. To stop the service:
```bash
docker compose down
```
  ## Error Handling
  
  The API returns appropriate HTTP status codes:
  
  - **200 OK** - Successful operation
  - **400 Bad Request** - Invalid JSON or empty key
  - **404 Not Found** - Key does not exist
  - **415 Unsupported Media Type** - Content-Type is not application/json
  - **500 Internal Server Error** - Failed to persist data

## Usage examples

### Store Data (PUT) :

Store a key-value pair:

```bash
curl -i -X PUT http://localhost:8080/objects \
  -H "Content-Type: application/json" \
  -d '{
    "key": "user:1234",
    "value": {
      "name": "Amin Alavi",
      "age": 23,
      "email": "a.alavi@fum.ir"
    }
  }'
```

**Response:**
```
HTTP/1.1 200 OK
Date: Sat, 25 Oct 2025 10:42:53 GMT
Content-Length: 0
```

### Retrieve Data (GET) :

Retrieve a value by key:

```bash
curl -i http://localhost:8080/objects/user:1234
```

**Response (if key exists):**
```
HTTP/1.1 200 OK
Content-Type: application/json
Date: Sat, 25 Oct 2025 10:44:23 GMT
Content-Length: 84

{
      "name": "Amin Alavi",
      "age": 23,
      "email": "a.alavi@fum.ir"
    }
```

**Response (if key doesn't exist):**
```
HTTP/1.1 404 Not Found
Content-Type: text/plain; charset=utf-8
X-Content-Type-Options: nosniff
Date: Sat, 25 Oct 2025 10:45:28 GMT
Content-Length: 14

Key not found

```

### Try to get non-existent key :

Retrieve a non-existent key:

```bash
curl -i http://localhost:8080/objects/nonexistent
```

**Response:**
```
HTTP/1.1 404 Not Found
Content-Type: text/plain; charset=utf-8
X-Content-Type-Options: nosniff
Date: Sat, 25 Oct 2025 20:00:54 GMT
Content-Length: 14

Key not found
```

### Wrong Content-Type : 

Store a content-type other than Json:

```bash
curl -i -X PUT http://localhost:8080/objects   -H "Content-Type: text/plain"   -d '{"key": "test", "value": "data"}'
```

**Response:**
```
HTTP/1.1 415 Unsupported Media Type
Content-Type: text/plain; charset=utf-8
X-Content-Type-Options: nosniff
Date: Sat, 25 Oct 2025 19:51:28 GMT
Content-Length: 38

Content-Type must be application/json
```

### Empty key : 

Retrieve a value by empty key:

```bash
curl -i -X PUT http://localhost:8080/objects   -H "Content-Type: application/json"   -d '{"key": "", "value": "something"}'
```

**Response:**
```
HTTP/1.1 400 Bad Request
Content-Type: text/plain; charset=utf-8
X-Content-Type-Options: nosniff
Date: Sat, 25 Oct 2025 19:54:13 GMT
Content-Length: 20

Key cannot be empty
```

### Invalid JSON : 

Store invalid JSON:

```bash
curl -i -X PUT http://localhost:8080/objects   -H "Content-Type: application/json"   -d 'invalid json here'
```

**Response:**
```
HTTP/1.1 400 Bad Request
Content-Type: text/plain; charset=utf-8
X-Content-Type-Options: nosniff
Date: Sat, 25 Oct 2025 19:59:20 GMT
Content-Length: 20

Invalid JSON format
```
