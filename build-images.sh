#!/bin/bash

EXEC=docker

USER="xzhu0027"

TAG="latest"

for i in "frontend" "server-v1" "server-v2"
do
  IMAGE="echo-${i}-grpc-proxyless"
  DOCKERFILE="Dockerfile-${i}"
  
  echo Processing image ${image}
  $EXEC build -t "$USER"/"$IMAGE":"$TAG" -f "$DOCKERFILE" .
  $EXEC push "$USER"/"$IMAGE":"$TAG"

  echo
done
