FROM golang:1.19-alpine3.15
RUN mkdir /app
ADD . /app
WORKDIR /app

# Copia el archivo .env a la carpeta de trabajo en la imagen
COPY .env /app/.env

RUN go mod download && go build -o main ./main.go
CMD /app/main
