FROM golang:1.19-alpine3.15
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go mod download && go build -o main ./main.go
CMD /app/main