FROM golang:1.17-alpine as builder

# first (build) stage
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 go build -o backend-server

# final (target) stage
FROM scratch
COPY --from=builder /app/backend-server .
CMD ["./backend-server"]
