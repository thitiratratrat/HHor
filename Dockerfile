FROM golang:1.17-buster AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
ENV GO111MODULE on
RUN go mod download

COPY . .

RUN go build -o /hhor 


##
## Deploy
##
FROM debian:buster-slim

ARG DEBIAN_FRONTEND=noninteractive

WORKDIR /app

COPY --from=build /hhor /app/hhor
COPY ./template /app/template

EXPOSE $PORT

ENTRYPOINT ["/app/hhor"]