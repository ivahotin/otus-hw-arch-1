## Build
FROM golang:1.16-buster AS build

WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY *.go .
RUN go build -o /server

## Deploy
FROM gcr.io/distroless/base-debian10
WORKDIR /
COPY --from=build /server /server
EXPOSE 8000
USER nonroot:nonroot

ENTRYPOINT ["/server"]