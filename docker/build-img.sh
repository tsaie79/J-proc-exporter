#!/bin/bash

docker build -t jlabtsai/process-exporter:pgid --no-cache .
docker push jlabtsai/process-exporter:pgid