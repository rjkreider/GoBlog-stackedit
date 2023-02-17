# GoBlog-stackedit
Integrating stackedit markdown editor in GoBlog

# Installation

Navigate into your GoBlog directory.

Clone the repository

```bash
git clone https://github.com/rjkreider/GoBlog-stackedit plugins/stackedit
```

Update the GoBlog configuration in ./conf/config.yml

```yaml
plugins:
  - path: ./plugins/stackedit
    import: stackedit
```

Update the docker configuration to make sure the plugins volume is mapped.

```yaml
        volumes:
            - ./config:/app/config # Config directory
            - ./data:/app/data # Data directory, used for database, keys, upload
s etc.
            - ./static:/app/static # Static directory, if you want to publish st
atic files
            - ./plugins:/app/plugins
```

Start the GoBlog container, make sure no errors.

```bash
docker compose up
```

If no errors, then start it up in daemon.

```bash
docker compose up -d
```

# Notes

This is under significant development and may break things.  You can comment out the ./conf/config.yaml plugin lines if you want to disable the plugin if it breaks your GoBlog. 
