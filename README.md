# rabbitmq-publish-cron

Simple docker container to publish a fixed message to a specified queue. Created to be used with k8s CRON scheduling.

## Local Docker Build & Run

```
docker build -t rabbitmq-publish-cron:latest .
docker run --env-file=.env rabbitmq-publish-cron:latest
```
