# Features in this branch
This branch requires users to place your own config file `$HOME/config.yml`. Template of config.yml is shown below, and you can also find it in [process-exporter](https://github.com/ncabatoff/process-exporter.git) repository.

## Template of config.yml
```yaml
process_names:
- name: "{{.ExeBase}}:{{.Username}}"
  comm: