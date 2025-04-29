FROM golang:1.2-alpine
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o main ./cmd/server
EXPOSE 8080
CMD [ "./main" ]