FROM golang:1.20 as builder

WORKDIR /app

COPY . .

RUN go env -w  GOPROXY=https://goproxy.cn,direct && go mod tidy && \
    CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o main main.go 

FROM busybox

COPY --from=builder /app/main .

EXPOSE 3000/tcp

CMD ["./main"]