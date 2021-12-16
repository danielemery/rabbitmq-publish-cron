FROM golang:1.16 as build
COPY . /app/
WORKDIR /app
RUN go install
RUN CGO_ENABLED=0 GOOS=linux go build

FROM scratch
WORKDIR /app
COPY --from=build /app/rabbitmq-publish-cron /app/rabbitmq-publish-cron
ENTRYPOINT [ "/app/rabbitmq-publish-cron" ]
