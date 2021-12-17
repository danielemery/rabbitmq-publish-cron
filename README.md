# rabbitmq-publish-cron

Simple docker container to publish a fixed message to a specified queue. Created to be used with k8s CRON scheduling.

## Usage

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

### Helm Publish

```sh
# If not already installed, install the helm s3 plugin
helm plugin install https://github.com/hypnoglow/helm-s3.git
# Add the helm repo
helm repo add demery-helm s3://helm.demery.com.au
# List to double check
helm repo list
# Package up the chart ready for deployment
helm package ./helm
# Push the bundle to the s3 repo (replace file with newly generated)
helm s3 push ./rabbitmq-publish-cron-0.1.0.tgz demery-helm
```

Note:
For the very first publish the following was run:

```sh
helm S3 init s3://helm.demery.com.au
```
