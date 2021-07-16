FROM golang:buster as build

WORKDIR /go/src/greeter-server

RUN curl -o main.go https://github.com/grpc/grpc-go/blob/91e0aeb192456225adf27966d04ada4cf8599915/examples/features/reflection/server/main.go && \
  go mod init greeter-server && \
  go mod tidy && \
  go build -o /greeter-server main.go

FROM gcr.io/distroless/base-debian10

COPY --from=build /greeter-server /

EXPOSE 50051

CMD ["/greeter-server"]
