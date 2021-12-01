#builds and run test
FROM golang:1.16-buster AS build
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY *.go ./
#Running tests
RUN go test -v
RUN go build -o /client

FROM gcr.io/distroless/base-debian10
WORKDIR /
COPY --from=build /client /client
USER nonroot:nonroot
ENTRYPOINT ["/client"]
CMD ["/short/ab"]