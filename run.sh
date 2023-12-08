#!/bin/bash


# Call the binary with the --pgid flag and set the path of proc to /host_proc
/usr/local/bin/gen-config --pgid $PGID_FILE --procpath /host_proc --outyml config.yml

# Run the process-exporter command
/usr/local/bin/process-exporter --procfs /host_proc --config.path config.yml --web.listen-address=:$PROCESS_EXPORTER_PORT