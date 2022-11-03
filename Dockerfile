FROM golang:1.19.3 AS build

WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . ./
RUN go build -o /miniserver

## Deploy
FROM gcr.io/distroless/base-debian11
WORKDIR /
COPY --from=build /miniserver /miniserver
EXPOSE 5508
USER nonroot:nonroot
VOLUME [ "/public" ]
ENTRYPOINT ["/miniserver"]