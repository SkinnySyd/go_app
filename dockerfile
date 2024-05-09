# Build the application
FROM golang:1.22.2 AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go  .
COPY . .


# COPY templates /app/templates

RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-hello

# Deploy the application binary into a lean image
FROM gcr.io/distroless/base-debian11 AS release

WORKDIR /

COPY --from=build /docker-hello /docker-hello
# COPY --from=build /app/templates /templates


EXPOSE 8088

USER nonroot:nonroot

ENTRYPOINT ["/docker-hello"]