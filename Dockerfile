FROM golang:1.16-alpine3.12 AS builder
LABEL stage=builder
WORKDIR /app
COPY . /app
RUN cd /app && go build -o main .

FROM adoptopenjdk/openjdk13:alpine-jre
RUN apk --no-cache add ca-certificates
RUN wget https://github.com/AsamK/signal-cli/releases/download/v0.8.1/signal-cli-0.8.1.tar.gz \
    && tar xf signal-cli-0.8.1.tar.gz -C /opt \
    && ln -sf /opt/signal-cli-0.8.1/bin/signal-cli /usr/local/bin/
COPY --from=builder /app/main .
COPY /configs /configs
ENTRYPOINT ./main
