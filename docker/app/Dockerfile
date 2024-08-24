FROM golang:1.23-alpine AS builder

WORKDIR /build
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .

RUN go build -o /service ./cmd/service/main.go

FROM golang:1.23-alpine AS service
COPY . .
WORKDIR /app
ENTRYPOINT [ "/app/docker/app/entrypoint.sh" ]


FROM golang:1.23-alpine AS worker
COPY --from=builder service /bin/service