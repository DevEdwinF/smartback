# FROM golang:1.19-alpine3.15
# RUN mkdir /app
# ADD . /app
# WORKDIR /app

# COPY .env /app/.env

# RUN go mod download && go build -o main ./main.go
# CMD /app/main

FROM golang:1.19-alpine3.15

WORKDIR /app

COPY .env /app/.env

COPY . /app

RUN go mod download && go build -o main ./main.go

CMD ["./main"]
