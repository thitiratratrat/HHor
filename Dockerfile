FROM golang:1.17-buster AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
ENV GO111MODULE on
RUN go mod download

COPY *.go ./
COPY /src ./src

RUN go build -o /hhor 


##
## Deploy
##
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /hhor /hhor

EXPOSE 8081

USER nonroot:nonroot

ENTRYPOINT ["/hhor"]