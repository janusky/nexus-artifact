# Demo

Have a nexus server certificate ([Copy Certificate](#copy-certificate-file)).

> NOTE: Run Nexus from docker image or use existing.
> * [Instancy Nexus](#instancy-nexus)

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

## Instancy Nexus

Run the image and enter Nexus to create repository `demo-raw`.

* <https://hub.docker.com/r/bradbeck/nexus-https>
* <https://github.com/bradbeck/nexus-https>

```sh
# Install
docker pull bradbeck/nexus-https

# Run
docker run -p 8443:8443 bradbeck/nexus-https

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

### Copy certificate file

Nexus must be running to copy the certificate.

```sh
# cd nexus-artifact/demo (cd .)

# Format PEM
echo | \
    openssl s_client -servername localhost -connect localhost:8443 2>/dev/null | \
    openssl x509 -outform PEM > nexus.server.pem

# Used to copy into `nexus.certsDefault`
# Print Certificate
echo | \
    openssl s_client -servername localhost -connect localhost:8443 2>/dev/null | \
    openssl x509 -text
```