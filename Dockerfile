FROM golang:latest
RUN mkdir /app
COPY . /app
WORKDIR /app
RUN go build -o server ./cmd/main.go
EXPOSE 8000
CMD [ "/app/server" ]
