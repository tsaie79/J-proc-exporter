#!/bin/bash

# Environment Variables:
# GROUP_FOR_CONFIG: This variable is used to group processes in the generated config file. If not set, it defaults to an empty string.
# CMDS_FOR_CONFIG: This variable is used to specify the commands for which the config file should be generated. If not set, the PGID_FILE is used instead.
# PGID_FILE: This variable is used to specify a file containing the PGID for which the config file should be generated. This is used only if CMDS_FOR_CONFIG is not set.
# PROCESS_EXPORTER_PORT: This variable is used to specify the port on which the process-exporter should listen.

#if $GROUP is set, then export GROUP to $GROUP, else export GROUP to ""
if [ -z "$GROUP_FOR_CONFIG" ]; then
    echo "GROUP is empty"
    export GROUP_FOR_CONFIG=""
else
    echo "GROUP is NOT empty"
fi



# if $CMD_FOR_CONFIG is set, then use it to generate the config.yml
if [ -z "$CMDS_FOR_CONFIG" ]; then
    echo "CMDS_FOR_CONFIG is empty"
    /usr/local/bin/gen-config --pgid "$PGID_FILE" --procpath /host_proc --outyml config.yml --group "$GROUP_FOR_CONFIG"
else
    echo "CMDS_FOR_CONFIG is NOT empty"
    /usr/local/bin/gen-config --cmds "$CMDS_FOR_CONFIG" --procpath /host_proc --outyml config.yml --group "$GROUP_FOR_CONFIG"
fi

# print the config.yml
echo "config.yml:"
cat config.yml

# Run the process-exporter command
/usr/local/bin/process-exporter --procfs /host_proc --config.path config.yml --web.listen-address=:"$PROCESS_EXPORTER_PORT"