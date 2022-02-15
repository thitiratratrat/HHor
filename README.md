# HHor

## Deployed url
https://hhor-fydi7fajxq-uc.a.run.app/

## Swagger documentation
https://hhor-fydi7fajxq-uc.a.run.app/swagger/index.html

<br/>

## Run locally
### Prerequisites
- Docker

### Run
Build and run docker image
```
docker compose up --build
```

<br/>

## Set up for development
### Prerequisites
- Go version at least 1.13

### Set up
1. Enable Go module in go env
```
export GO111MODULE=on
```

2. Build and run development postgres image
```
docker build -t hhor-db -f Dockerfile.dev .
docker run -p 5432:5432 hhor-db
```

3. Set environment variables
```
export PORT=8080
export INSTANCE_CONNECTION_NAME=localhost
export DB_USER=postgres
export DB_PASS=postgres
export DB_NAME=hhor
export DB_PORT=5432
```

4. Download dependencies
```
go mod download
```

4. Run code
```
go run .
```
or 
```
go run main.go
```