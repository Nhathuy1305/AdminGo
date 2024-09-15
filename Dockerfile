FROM golang:1.23-alpine AS build

WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o run .

FROM alpine
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=build /build/run .
EXPOSE 8000
CMD ["/app/run"]

#https://github.com/miguno/golang-docker-build-tutorial/tree/main