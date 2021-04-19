FROM golang:1.16-alpine3.13 AS builder
LABEL stage=builder
WORKDIR /app
COPY . /app
RUN cd /app && go build -o main .

FROM alpine:3.13
RUN apk --no-cache add ca-certificates
COPY --from=builder /app/main .
COPY /configs /configs/
ENTRYPOINT ./main
