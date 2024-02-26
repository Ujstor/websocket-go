FROM golang:1.22-alpine as build 

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o server cmd/api/main.go
RUN go build -o client-chat client/ws-chat/main.go
RUN go build -o client-data client/ws-data-read/main.go


FROM alpine:3.19.0 as server
WORKDIR /app
COPY --from=build /app/server /app/server
COPY --from=build /app/.env /app/.env
EXPOSE ${PORT}
CMD ["./server"]


FROM alpine:3.19.0 as chat
WORKDIR /app
COPY --from=build /app/client-chat /app/client-chat
CMD ["./client-chat"]

FROM alpine:3.19.0 as data
WORKDIR /app
COPY --from=build /app/client-data /app/client-data
CMD ["./client-data"]