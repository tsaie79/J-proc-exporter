#!/bin/bash

docker run -d --net=host --user "$(id -u)" --volume "/workspaces/test-ersap-wf/prom/grafana-data:/var/lib/grafana" grafana/grafana-enterprise