FROM golang:1.19 AS base
WORKDIR /app
COPY . .
FROM base AS build-server
RUN go build -o /bin/server ./main.go


# FROM scratch AS Server
# COPY --from=build-server /app/server /server
ENTRYPOINT [ "/bin/server" ]




