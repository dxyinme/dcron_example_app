# dcron_example_app

## Design Doc

[design-doc](documentation/design-doc.md)

## CI 

[![Docker Image CI](https://github.com/dxyinme/dcron_example_app/actions/workflows/docker-image.yml/badge.svg?branch=main)](https://github.com/dxyinme/dcron_example_app/actions/workflows/docker-image.yml)

## Build

```bash
# swagger gen
cd app
swag init
```

```bash
# build
cd app
go build 
```

```bash
# docker build
docker build -f deployment/Dockerfile -t dcron_example_app:latest .
```

## Run in docker compose

You can change the app.yaml volume to change the configuration.
```yaml
version: '3.0'
services:
  app:
    image: dcron_example_app:latest
    scale: 3
    volumes:
      - /path/to/yourself-app.yaml:/app/etc/app.yaml
    ports:
      - 10010-10012:8080
```

## Features

 status | Features
---|---
Done|MySQL
TODO|PostgreSQL
Done|Monitor
TODO|Use ETCD to run dcron
TODO|Use service discovery to implement inner call.

