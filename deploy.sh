#!/bin/bash -x

gcloud functions deploy PubSubHandlerForCloudBuild \
  --allow-unauthenticated \
  --trigger-topic cloud-builds \
  --runtime go111 \
  --set-env-vars "GCSN_SLACK_URL=${GCSN_SLACK_URL},GCSN_SLACK_CHANNEL=${GCSN_SLACK_CHANNEL},GCSN_START_NOTIFY_DISABLE=false,GCSN_END_NOTIFY_DISABLE=false"
