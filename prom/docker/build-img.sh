#!/bin/bash

img_name="jlabtsai/process-exporter:v1.1"
docker build -t $img_name --no-cache .
docker push $img_name