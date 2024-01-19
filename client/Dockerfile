FROM golang:1.19 AS base

WORKDIR /app

COPY go.mod .

COPY main.go .
FROM base as build-client
RUN go build -o /bin/client main.go
ENTRYPOINT [ "/bin/client" ]

# FROM scratch AS client
# COPY --from=build-client /bin/client client
# ENTRYPOINT [ "/bin/client" ]