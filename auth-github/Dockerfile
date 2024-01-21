FROM golang:1.20 as builder

WORKDIR /app

COPY . .

# go 交叉编译: amd64
RUN go env -w  GOPROXY=https://goproxy.cn,direct && go mod tidy && \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main main.go 

RUN chmod +x /app/main

FROM alpine:3.19

#RUN apk add --update bash

COPY --from=builder /app/main .

EXPOSE 3000/tcp

CMD ["./main"]