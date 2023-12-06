#!/bin/bash

/get_cmds.sh $PGID_FILE
/bin/process-exporter --procfs /host_proc --config.path $HOME/config.yml --web.listen-address=:$PROCESS_EXPORTER_PORT
