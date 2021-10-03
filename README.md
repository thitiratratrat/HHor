# HHor

## Swagger Documentation
Head to route /swagger/index.html#/

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