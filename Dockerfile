FROM golang:alpine

COPY . /app
WORKDIR /app

RUN go build -o main main.go

EXPOSE 3030

ENTRYPOINT [ "/app/main" ]
