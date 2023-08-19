# dcron_example_app

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