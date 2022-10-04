FROM golang:1.14-alpine as build

WORKDIR /app

RUN apk update
RUN apk add gcc git g++ net-tools

COPY go.mod ./
RUN go mod download
RUN go mod tidy

COPY ./ ./

FROM build as grpcserver
WORKDIR /app

COPY --from=build /app .
RUN go build -o server -a ./cmd/grpc
EXPOSE 8080
CMD ["./server"]


FROM build as httpserver
WORKDIR /app

COPY --from=build /app .
RUN go build -o server -a ./cmd/http
EXPOSE 8081
CMD ["./server"]

FROM build as clientserver
WORKDIR /app

COPY --from=build /app .
RUN go build -o server -a ./cmd/client
EXPOSE 8081
CMD ["./server"]
