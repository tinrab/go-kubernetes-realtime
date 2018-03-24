#!/bin/bash

docker build -t local/echo -f Dockerfile.echo .

kubectl apply -f kube-echo.yaml
