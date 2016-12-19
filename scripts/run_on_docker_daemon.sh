#!/bin/bash

docker run -d \
  --name $1 \
  -e SAKURAIOT_AUTH_TOKEN \
  -e SAKURAIOT_AUTH_SECRET \
  -e SAKURAIOT_MODULE_ID \
  -e SAKURAIOT_MODULE_PASS \
  -e SAKURAIOT_MODULE_SECRET \
  $1:latest ${@:2}