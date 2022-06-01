# Using this repo

## Get the `pack` CLI
Visit the [docs](https://buildpacks.io/docs/tools/pack/) site and get it for your platform.

## Build

### Heroku
```shell
pack build <image-name> --path <path-to-source-code> --builder heroku/buildpacks:20
```

### Google
```shell
pack build <image-name> --path <path-to-source-code> --builder gcr.io/buildpacks/builder:v1
```

### Paketo
```shell
pack build <image-name> --path <path-to-source-code> --builder paketobuildpacks/builder:base
```

### Docker: NodeJS
```shell
docker build -t <image-name> -f Dockergfile.node nodejs
```

### Docker: Go
```shell
docker build -t <image-name> -f Dockergfile.go go
```

## Analyze
Analyze the image for size, for Software Bill Of Material with `docker sbom` and for vulnerabilities with `docker scan`:
```shell
sh ./analyze.sh <image-name>
```