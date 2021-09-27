FROM golang:1.17-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
ENV GO111MODULE on
RUN go mod download

COPY *.go ./
COPY /src ./src

RUN go build -o ./hhor 

EXPOSE 8081

CMD [ "./hhor" ]