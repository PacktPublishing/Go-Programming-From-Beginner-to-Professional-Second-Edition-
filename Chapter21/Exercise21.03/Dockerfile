FROM golang:latest AS builder 
ENV CGO_ENABLED=0 
WORKDIR /app 
COPY go.mod go.sum ./ 
COPY main.go ./ 
RUN go mod download 
RUN go build -o monitored_app .
FROM scratch 
COPY --from=builder /app/monitored_app /. 
EXPOSE 8080 
CMD ["./monitored_app"] 