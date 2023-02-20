# nexus-artifact

Application to perform PUT and GET of artifacts in Nexus repository. 

## Objective

Run upload and get of artifacts in Nexus repository.

## Build

Build application executable

```sh
# Build
go build -ldflags "-X 'main.version=dev' -X 'main.date=$(date '+%Y%m%d%H%M')' -X 'main.commit=$(git rev-parse HEAD)'"

# Show version
./nexus-artifact -v
```

## Run

There should be a [file list](demo/DEMO.md#create-files-to-work), which represent the artifacts that you want to recover or upload.

> NOTE: You must have an application executable.
> * [Build App](#build)

```sh
# GET
./nexus-artifact \
  --mode "GET" \
  --list "./artifacts.json" \
  --trust "./nexus.server.pem" \
  --parallelism 33

# PUT
./nexus-artifact \
  --mode "PUT" \
  --list "./artifacts.json" \
  --trust "./nexus.server.pem" \
  --parallelism 10 \
  --rerun-n 2 \
  --retry-failures
``` 

### JSON artifacts file

PUT -> The `JSON` data model requires indicating **url** and **file**. Optionally **contenttype**.

```json
{
  "artifacts": [
    {
      "url": "https://localhost:8443/repository/demo-raw/nexus-artifact/1.0.0/nexus-artifact",
      "file": "./demo/nexus-artifact",
      "contenttype": "text/plain"
    }
  ]
}
```

GET -> The `JSON` data model requires indicating **url**. Not required **file** and **contenttype**.

```json
{
  "artifacts": [
    {
      "url": "https://localhost:8443/repository/demo-raw/nexus-artifact/1.0.0/nexus-artifact"
    }
  ]
}
```

## Documents

[DEMO](./demo/DEMO.md)

[TODO](./TODO.md)