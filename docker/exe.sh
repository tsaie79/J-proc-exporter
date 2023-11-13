#!/bin/bash

#repeat launch every 3 seconds and kill the launch process
while true; do
    setsid /run-exporter.sh& echo $! > $HOME/launch-exporter.sid
    sleep 10
    kill -9 `pgrep -s $(cat $HOME/launch-exporter.sid)`
done


## This has a problem that when the process-exporter is killed, the ssh tunnel will fail