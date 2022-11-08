## Build
FROM golang:1.19-buster AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY cmd/ ./cmd
COPY pkg/ ./pkg

WORKDIR cmd/
RUN go build -o /dist

## Deploy
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /dist /dist

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/dist"]