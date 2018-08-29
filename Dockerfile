FROM arm32v6/golang:alpine as builder

WORKDIR /go/src/app
COPY . .

RUN apk add --no-cache binutils && \
    CGO_ENABLED=0 GOOS=linux go build -a -tags netgo -ldflags '-w' -o collatz && \
    strip collatz


FROM scratch

WORKDIR /collatz/
EXPOSE 80
ENTRYPOINT ["./collatz"]

COPY --from=builder /go/src/app/collatz .
COPY ./static/ ./static/
COPY ./templates/ ./templates/
