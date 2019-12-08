#!/bin/bash -x

gcloud functions deploy PubSubHandlerForCloudBuild \
  --trigger-topic cloud-builds \
  --runtime go111 \
  --set-env-vars "GCSN_SLACK_URL=${GCSN_SLACK_URL}"
