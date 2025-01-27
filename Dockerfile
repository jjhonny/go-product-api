FROM golang:1.22.3 as stage1

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main cmd/main.go

FROM scratch

COPY --from=stage1 /app/main /

ENTRYPOINT ["/main"]