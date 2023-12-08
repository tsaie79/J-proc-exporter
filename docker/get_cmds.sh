#!/bin/bash

# This give a list of pids that are in the container. I want to loop get the result of the cmd
#cat /proc/pid/stat | awk '{print $2}' and remove the perenthesis

cp /config.yml $PROCESS_EXPORTER_CONFIG

for pid in $(pgrep -g `cat $PGID_FILE`);
do
    export cmd=$(cat /proc/$pid/stat | awk '{print $2}' | sed 's/(//g' | sed 's/)//g')
    echo $cmd

    sed -i "s/comm:/&\n  - \"$cmd\"/" $PROCESS_EXPORTER_CONFIG
done

