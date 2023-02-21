## nexus-artifact demo



## Demo

Requirements for execution if necessary

* **1** Have a nexus server certificate ([Copy Certificate](#copy-certificate-file)).
* **2** Run Nexus from [docker image](#instancy-nexus) or use existing.

There should be a [file list](#create-files-to-work), which represent the artifacts that you want to recover or upload.

>If fail
>```sh
>error fail run 1: [Put "https://localhost:8443/repository/demo-raw/nexus-artifact/1.0.0/nexus-artifact.sha1": x509: certificate relies on legacy Common Name field, use SANs or temporarily enable Common Name matching with GODEBUG=x509ignoreCN=0]
>```
>Solve
>```sh
>export GODEBUG=x509ignoreCN=0
>```

```sh
# Access path
cd ~/nexus-artifact/demo

# Put
./nexus-artifact \
  --mode "PUT" \
  --list "./artifacts.json" \
  --trust "./nexus.server.pem" \
  --debug

# Get
./nexus-artifact \
  --mode "GET" \
  --list "./artifacts.json" \
  --trust "./nexus.server.pem" \
  --parallelism 10

# NOTA: In the const.go file the certsDefault attribute must have the appropriate value
#
# Run until that put all (--trust default)
./nexus-artifact \
  --mode "PUT" \
  --list "./artifacts.json" \
  --parallelism 10 \
  --rerun-n 5 \
  --retry-failures

./nexus-artifact \
  --mode "GET" \
  --list "./artifacts.json" \
  --parallelism 10 \
  --rerun-n 5 \
  --retry-failures
```

Open in Browser

* https://localhost:8443/service/rest/repository/browse/demo-raw/

## Instancy Nexus

Run the image and enter Nexus to create repository `demo-raw`.

* <https://hub.docker.com/r/bradbeck/nexus-https>
* <https://github.com/bradbeck/nexus-https>

```sh
# Donwload Nexus Image
docker pull bradbeck/nexus-https:v3.38.1

# Run
docker run -p 8443:8443 bradbeck/nexus-https:v3.38.1

# Test
curl -kvu admin:admin123 https://localhost:8443/service/metrics/ping
```

Access https://localhost:8443/

> IMPORTANT (Initial config)  
> Your admin user password is located in /nexus-data/admin.password on the server.

```sh
# Read container identifier
# docker ps -aqf "ancestor=bradbeck/nexus-https
docker exec -ti $(docker ps -aqf "ancestor=bradbeck/nexus-https") sh -c "cat /nexus-data/admin.password"
```

> Indicate new password `admin123`  
> Enable anonymous access  
> create repository `demo-raw` type **raw** (default config) 

## Create files to work

Use the `nexus-artifact` file or create your own version (Start compile).

> Start compile
>>```sh
>>cd ~/nexus-artifact
>>go build -ldflags="-X main.version='demo' -X 'main.date=$(date +%Y/%m/%d-%H:%M)'"
>>``` 

```sh
# cd ~/nexus-artifact/demo (cd .)
export nexusRepo=https://localhost:8443/repository/demo-raw/
export fileUp=nexus-artifact
export artefact=nexus-artifact/1.0.0/$fileUp 

# MD5 y SHA1 (md5sum $file | cut -d' ' -f1 > $file_out.md5)
md5sum $fileUp > $fileUp.md5
sha1sum $fileUp > $fileUp.sha1

cat > artifacts.json <<EOF
{
  "artifacts": [
    {
      "url": "$nexusRepo$artefact",
      "file": "./$fileUp"
    },
    {
      "url": "$nexusRepo$artefact.md5",
      "file": "./$fileUp.md5"
    },
    {
      "url": "$nexusRepo$artefact.sha1",
      "file": "./$fileUp.sha1",
      "contenttype": "text/plain"
    }
  ]
}
EOF
```

## Copy certificate file

Nexus must be running to copy the certificate.

```sh
cd nexus-artifact/demo

# Format PEM
# -extfile <(printf "subjectAltName=DNS:localhost")
echo | \
    openssl s_client -servername localhost -connect localhost:8443 2>/dev/null | \
    openssl x509 -outform PEM > nexus.server.pem

# Used to set in the value of the `nexus.certsDefault` variable from the `nexus/const.go` file
# Print Certificate
echo | \
    openssl s_client -servername localhost -connect localhost:8443 2>/dev/null | \
    openssl x509 -text
```