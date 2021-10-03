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
1. Build docker image
```
docker build -t hhor:latest .
```

2. Run docker image
```
docker run -e PORT={VALID_PORT_NUMBER} -p {VALID_PORT_NUMBER}:{VALID_PORT_NUMBER} hhor:latest
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

2. Set port environment variable
```
export PORT={VALID_PORT_NUMBER}
```

3. Download dependencies
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