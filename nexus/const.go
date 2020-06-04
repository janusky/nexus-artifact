package nexus

// use only if you are sure it is indicated
const certsDefault = `
Certificate:
    Data:
        Version: 3 (0x2)
        Serial Number:
            0a:18:9e:cb:e2:e2:ee:ca:ed:4b:b6:14:4c:a6:fb:0e:ac:34:27:08
        Signature Algorithm: sha256WithRSAEncryption
        Issuer: CN = localhost
        Validity
            Not Before: Jun  3 21:38:18 2020 GMT
            Not After : Jul  3 21:38:18 2020 GMT
        Subject: CN = localhost
        Subject Public Key Info:
            Public Key Algorithm: rsaEncryption
                RSA Public-Key: (2048 bit)
                Modulus:
                    00:dd:ab:61:d7:d0:72:f0:49:38:d3:9c:a7:bc:97:
                    a8:f6:b1:7e:5c:4e:52:dc:b7:f5:22:b3:79:4a:81:
                    db:28:da:27:12:0a:e0:88:60:9e:ff:d8:d5:59:95:
                    cd:15:fe:2d:9d:b6:e1:56:f1:93:dd:9b:f5:90:b4:
                    51:27:1a:4a:6b:48:41:c9:cb:42:cb:74:c3:8d:66:
                    df:ad:3a:04:ba:82:d4:70:21:e4:82:a8:3f:89:71:
                    d9:09:3b:1b:18:f9:2f:dc:d7:01:98:81:fb:da:86:
                    0f:7b:9b:37:32:0a:05:e3:d1:d9:bd:e6:49:ab:38:
                    2d:3e:55:de:81:d7:37:b2:99:c4:bf:52:8a:18:8e:
                    f7:34:86:9c:8f:7f:a7:0f:50:9e:bf:ea:6f:30:ab:
                    79:28:25:71:8c:15:6e:52:af:94:78:6f:86:c6:66:
                    1c:09:ed:1d:ba:c4:98:6f:a3:64:22:48:ec:b4:f5:
                    6e:6b:0f:32:db:3f:c9:94:26:f2:6c:bc:b8:bd:e6:
                    34:63:cb:4a:8d:c3:9e:23:3b:36:06:84:bd:b6:a6:
                    71:6f:df:96:82:63:d6:a7:93:55:cb:74:33:9a:49:
                    7a:d0:72:62:8e:f4:2e:e7:85:96:6c:d8:12:ec:a9:
                    fc:00:42:7d:16:0f:96:9b:2a:a7:24:5e:74:a9:57:
                    a3:0d
                Exponent: 65537 (0x10001)
        X509v3 extensions:
            X509v3 Subject Key Identifier: 
                8E:48:3D:8A:6F:E1:E8:79:2A:65:AA:20:04:66:90:98:8A:F1:E4:C9
            X509v3 Authority Key Identifier: 
                keyid:8E:48:3D:8A:6F:E1:E8:79:2A:65:AA:20:04:66:90:98:8A:F1:E4:C9

            X509v3 Basic Constraints: critical
                CA:TRUE
    Signature Algorithm: sha256WithRSAEncryption
         a3:85:2b:f3:ee:e4:b2:36:02:db:7c:a3:85:ba:fc:89:21:b6:
         a6:ac:b9:78:59:c0:f6:cd:ae:b6:3b:03:6e:47:3b:be:aa:fd:
         78:50:7b:bc:18:35:9e:b1:d5:b9:53:69:62:41:d6:57:6e:e4:
         61:4b:a9:dc:49:d5:79:fe:d8:ef:43:03:77:49:78:08:4a:ed:
         02:40:33:79:a4:4f:f1:c0:4e:38:2c:39:ab:99:92:2c:d9:3b:
         6d:80:2b:a1:2c:4b:05:02:9f:6f:fe:b1:73:f6:a0:c4:a1:89:
         d6:8a:82:43:53:db:35:e6:84:e5:8e:b6:7e:7d:dd:63:03:4e:
         63:a5:97:0c:d7:bc:72:3f:36:96:d2:17:91:9a:d5:c2:fd:b4:
         42:c3:e0:c1:35:d0:4a:bf:14:71:ac:6e:5f:56:60:d4:9f:76:
         77:d0:a8:53:f6:1d:7f:83:a4:69:d2:a3:1f:c2:68:04:d2:39:
         53:73:4c:e3:88:ed:42:7a:b4:8b:9b:90:96:9c:0d:62:91:c0:
         f5:92:76:f0:b9:44:1d:f5:fe:64:0d:3b:d9:75:6a:6b:3d:f3:
         fc:c3:95:d8:0f:53:6a:45:4b:ea:ef:1f:f7:4b:e1:25:6f:e2:
         f4:09:c3:f1:bf:db:02:35:d5:5f:03:6f:d9:8f:95:2a:5e:34:
         21:6c:88:c6
-----BEGIN CERTIFICATE-----
MIIDCTCCAfGgAwIBAgIUChiey+Li7srtS7YUTKb7Dqw0JwgwDQYJKoZIhvcNAQEL
BQAwFDESMBAGA1UEAwwJbG9jYWxob3N0MB4XDTIwMDYwMzIxMzgxOFoXDTIwMDcw
MzIxMzgxOFowFDESMBAGA1UEAwwJbG9jYWxob3N0MIIBIjANBgkqhkiG9w0BAQEF
AAOCAQ8AMIIBCgKCAQEA3ath19By8Ek405ynvJeo9rF+XE5S3Lf1IrN5SoHbKNon
EgrgiGCe/9jVWZXNFf4tnbbhVvGT3Zv1kLRRJxpKa0hByctCy3TDjWbfrToEuoLU
cCHkgqg/iXHZCTsbGPkv3NcBmIH72oYPe5s3MgoF49HZveZJqzgtPlXegdc3spnE
v1KKGI73NIacj3+nD1Cev+pvMKt5KCVxjBVuUq+UeG+GxmYcCe0dusSYb6NkIkjs
tPVuaw8y2z/JlCbybLy4veY0Y8tKjcOeIzs2BoS9tqZxb9+WgmPWp5NVy3Qzmkl6
0HJijvQu54WWbNgS7Kn8AEJ9Fg+WmyqnJF50qVejDQIDAQABo1MwUTAdBgNVHQ4E
FgQUjkg9im/h6HkqZaogBGaQmIrx5MkwHwYDVR0jBBgwFoAUjkg9im/h6HkqZaog
BGaQmIrx5MkwDwYDVR0TAQH/BAUwAwEB/zANBgkqhkiG9w0BAQsFAAOCAQEAo4Ur
8+7ksjYC23yjhbr8iSG2pqy5eFnA9s2utjsDbkc7vqr9eFB7vBg1nrHVuVNpYkHW
V27kYUup3EnVef7Y70MDd0l4CErtAkAzeaRP8cBOOCw5q5mSLNk7bYAroSxLBQKf
b/6xc/agxKGJ1oqCQ1PbNeaE5Y62fn3dYwNOY6WXDNe8cj82ltIXkZrVwv20QsPg
wTXQSr8UcaxuX1Zg1J92d9CoU/Ydf4OkadKjH8JoBNI5U3NM44jtQnq0i5uQlpwN
YpHA9ZJ28LlEHfX+ZA072XVqaz3z/MOV2A9TakVL6u8f90vhJW/i9AnD8b/bAjXV
XwNv2Y+VKl40IWyIxg==
-----END CERTIFICATE-----
`