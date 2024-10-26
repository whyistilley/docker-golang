FROM golang:1.23.2-bookworm AS build
ARG LDFLAGS=""
ARG TEST_LDFLAGS=""

WORKDIR /usr/src/app

COPY ./go.mod ./go.sum ./
RUN go mod download && go mod verify

COPY ./ ./
RUN go test -v -p 1 -race -failfast -ldflags="$TEST_LDFLAGS" ./...
RUN go build -v -p 1 -o /usr/src/app/bin/app -ldflags="$LDFLAGS" ./...

FROM debian:bookworm-slim

COPY --from=build /usr/src/app/bin/app /usr/local/bin/app

USER 65534

CMD ["/usr/local/bin/app"]
