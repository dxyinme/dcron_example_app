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
docker build -f deployment/Dockerfile -t dcronapp:latest .
```

## Features

 status | Features
---|---
Done|MySQL
TODO|PostgreSQL
TODO|Monitor
TODO|Use ETCD to 
TODO|Use service discovery to implement inner call.

