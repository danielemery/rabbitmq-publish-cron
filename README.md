# RabbitMQ Publish CRON

Simple docker container to publish a fixed message to a specified rabbitmq exchange. Created to be used as part of a Kubernetes `CronJob`.

This repository includes:

- A lightweight golang implementation of publishing a single message to a rabbitmq exchange.
- Docker packaging of the implementation.
- Helm chart packaging of the container running as a Kubernetes `CronJob`.

## Usage

### Run with helm (recommended)

#### Create and populate `values.yml` file.

| Name         | Description                                                                                                                                   | Default | Required |
| ------------ | --------------------------------------------------------------------------------------------------------------------------------------------- | ------- | -------- |
| rabbitURL    | Full `amqp://` url to connect to rabbit                                                                                                       | None    | Yes      |
| exchangeName | The name of the exchange to publish to, will be created if doesn't exist                                                                      | None    | Yes      |
| messageBody  | The message body to publish to the exchange                                                                                                   | None    | Yes      |
| cronSchedule | The schedule definition [See CRON Schedule Syntax](https://kubernetes.io/docs/concepts/workloads/controllers/cron-jobs/#cron-schedule-syntax) | None    | Yes      |

Example:

```yml
# values.yml
publish:
  rabbitURL: "amqp://local_test:local_test@localhost:5672/"
  exchangeName: "test_exchange"
  messageBody: "Hello, World!"
  cronSchedule: "*/1 * * * *"
```

#### Add repo and run chart

```sh
# Add https:// repo
helm repo add demery https://helm.demery.net
# Install rabbitmq-publish-cron into the `test` namespace
helm install publish-cron demery/rabbitmq-publish-cron -f values.yaml -n test
# Remove rabbitmq-publish-cron
helm uninstall publish-cron -n test
```

## Local Development

### Local Docker Build & Run

```sh
docker build -t rabbitmq-publish-cron:latest .
docker run --env-file=.env rabbitmq-publish-cron:latest
```

### Local Helm Run

```sh
helm install test_helm_deployment ./helm
```

## Manual Helm Publish

CircleCI is performing the publish of the helm chart to the helm repo on tag.
If for whatever reason you need to perform a manual publish of the chart to the helm repo, the following can be run:

```sh
# If not already installed, install the helm s3 plugin
helm plugin install https://github.com/hypnoglow/helm-s3.git
# Add the helm repo
helm repo add demery-s3 s3://helm.demery.net
# Package up the chart ready for deployment
helm package ./helm --version=0.0.2 --app-version=0.0.2
# Push the bundle to the s3 repo (replace file with newly generated)
helm s3 push --relative ./rabbitmq-publish-cron-0.0.2.tgz demery-s3
```

Note:
For the very first publish the following was run:

```sh
helm s3 init s3://helm.demery.net
```
