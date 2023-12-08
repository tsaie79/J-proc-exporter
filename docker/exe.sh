#!/bin/bash

/get_cmds.sh
/bin/process-exporter --procfs /host_proc --config.path $PROCESS_EXPORTER_CONFIG --web.listen-address=:$PROCESS_EXPORTER_PORT
