FROM golang:1.23-alpine AS build

WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o run .

FROM alpine
WORKDIR /app
COPY --from=build /build/run .
CMD ["/app/run"]