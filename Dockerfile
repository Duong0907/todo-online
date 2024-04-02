FROM golang:1.22.1-alpine3.19 AS build-stage

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . ./

RUN go build -o /todo-online

FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=build-stage /todo-online /todo-online

EXPOSE 8080

ENTRYPOINT ["/todo-online"]