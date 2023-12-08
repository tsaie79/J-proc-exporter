#!/bin/bash

docker build -t jlabtsai/process-exporter:pgid-go --no-cache .
docker push jlabtsai/process-exporter:pgid-go