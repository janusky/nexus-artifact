package nexus

// use only if you are sure it is indicated
//
// Print Certificate
//
// echo | openssl s_client -servername localhost -connect localhost:8443 2>/dev/null | openssl x509 -text
const certsDefault = `
Certificate:
    Data:
        Version: 3 (0x2)
        Serial Number:
            5f:42:2c:f1:88:c2:c6:d0:0a:39:f1:83:28:44:8b:c7:92:80:48:22
        Signature Algorithm: sha256WithRSAEncryption
        Issuer: CN = localhost
        Validity
            Not Before: Feb 20 20:28:20 2023 GMT
            Not After : Mar 22 20:28:20 2023 GMT
        Subject: CN = localhost
        Subject Public Key Info:
            Public Key Algorithm: rsaEncryption
                RSA Public-Key: (2048 bit)
                Modulus:
                    00:ce:ee:23:61:76:e9:0f:ab:a8:7f:d5:df:f4:19:
                    a5:f7:7d:6f:5b:fb:e3:79:d3:d6:a3:07:7b:7d:34:
                    31:00:2b:9b:cf:c3:ee:2d:ab:0f:94:c5:1d:d1:8a:
                    34:48:39:da:d0:8c:9b:a8:78:1d:2c:e6:f8:93:8d:
                    98:f2:16:4a:41:16:a8:0e:04:f6:09:46:00:5e:bb:
                    13:66:fe:3a:38:3c:62:db:fd:20:c0:de:40:4e:60:
                    71:e2:6b:c4:e7:9c:d2:d9:bd:97:26:b0:0d:61:3a:
                    35:c3:f8:62:9a:58:ce:51:48:06:52:0e:74:cd:0a:
                    2a:41:8d:77:c9:0f:e9:20:05:3b:a2:ad:65:a0:92:
                    92:81:0a:8f:55:06:c5:8b:a6:aa:36:06:fc:bc:a4:
                    38:f8:d0:5a:c8:79:e3:9c:a5:a0:9e:05:4e:4a:c0:
                    fd:12:2a:00:06:d6:21:cb:8f:1a:52:f6:53:3c:42:
                    90:81:97:cb:b0:38:bd:e1:47:3c:45:44:6e:fe:ab:
                    80:a0:7e:a9:39:f1:8e:5d:04:12:e0:78:d3:ac:59:
                    2a:81:c4:cb:17:42:96:d1:c1:2d:33:c8:18:6b:57:
                    7c:f5:52:05:2c:94:70:0b:1f:9e:3d:75:c6:01:75:
                    e6:b6:6d:95:65:e4:c5:0a:83:77:33:45:51:a1:f3:
                    96:7d
                Exponent: 65537 (0x10001)
        X509v3 extensions:
            X509v3 Subject Key Identifier:
                80:32:94:62:D8:A5:03:E7:6C:E1:E3:3F:8E:CF:70:CF:18:F9:22:0E
            X509v3 Authority Key Identifier:
                keyid:80:32:94:62:D8:A5:03:E7:6C:E1:E3:3F:8E:CF:70:CF:18:F9:22:0E

            X509v3 Basic Constraints: critical
                CA:TRUE
    Signature Algorithm: sha256WithRSAEncryption
         99:44:d9:05:21:91:e9:47:4f:a3:ff:eb:c0:76:48:ae:88:24:
         c1:a4:c0:9c:91:fe:5c:e1:79:85:21:17:f5:84:66:28:9f:c2:
         16:ae:ab:ea:64:3e:09:94:86:2f:0f:7a:7f:d0:60:90:f3:1a:
         d4:c7:32:0c:b6:f7:a7:f1:be:04:a8:bc:36:ba:c2:55:5a:fa:
         52:49:c8:7e:c6:f4:21:ba:ff:32:56:2a:ad:9d:e5:77:cd:27:
         f8:72:c5:f9:c6:7c:dc:31:40:3e:07:39:b0:99:e8:49:e8:de:
         84:d4:9d:2d:77:7f:d9:3e:02:ac:a1:67:7b:07:ab:71:82:81:
         c4:0a:e7:9a:c4:63:a1:7a:77:0d:c2:9d:8e:c9:ec:89:f4:91:
         b9:8d:01:91:25:ba:82:cd:0c:d8:85:23:fa:32:50:39:b6:64:
         56:85:2f:38:54:a6:28:b2:88:96:94:83:94:e2:aa:b8:02:a3:
         e1:b7:54:21:17:0a:19:38:eb:c9:86:02:9d:2f:60:b8:f4:09:
         2e:cf:6e:ac:30:3b:b0:39:2d:c4:0f:79:ce:2f:50:c6:80:7d:
         a0:2c:05:4f:08:39:0c:2e:a7:66:9d:ca:23:09:cf:26:a9:6a:
         77:a5:72:e3:39:5f:59:94:e1:f3:ec:6f:0e:7b:4d:d2:78:f7:
         1f:fa:87:ac
-----BEGIN CERTIFICATE-----
MIIDCTCCAfGgAwIBAgIUX0Is8YjCxtAKOfGDKESLx5KASCIwDQYJKoZIhvcNAQEL
BQAwFDESMBAGA1UEAwwJbG9jYWxob3N0MB4XDTIzMDIyMDIwMjgyMFoXDTIzMDMy
MjIwMjgyMFowFDESMBAGA1UEAwwJbG9jYWxob3N0MIIBIjANBgkqhkiG9w0BAQEF
AAOCAQ8AMIIBCgKCAQEAzu4jYXbpD6uof9Xf9Bml931vW/vjedPWowd7fTQxACub
z8PuLasPlMUd0Yo0SDna0IybqHgdLOb4k42Y8hZKQRaoDgT2CUYAXrsTZv46ODxi
2/0gwN5ATmBx4mvE55zS2b2XJrANYTo1w/himljOUUgGUg50zQoqQY13yQ/pIAU7
oq1loJKSgQqPVQbFi6aqNgb8vKQ4+NBayHnjnKWgngVOSsD9EioABtYhy48aUvZT
PEKQgZfLsDi94Uc8RURu/quAoH6pOfGOXQQS4HjTrFkqgcTLF0KW0cEtM8gYa1d8
9VIFLJRwCx+ePXXGAXXmtm2VZeTFCoN3M0VRofOWfQIDAQABo1MwUTAdBgNVHQ4E
FgQUgDKUYtilA+ds4eM/js9wzxj5Ig4wHwYDVR0jBBgwFoAUgDKUYtilA+ds4eM/
js9wzxj5Ig4wDwYDVR0TAQH/BAUwAwEB/zANBgkqhkiG9w0BAQsFAAOCAQEAmUTZ
BSGR6UdPo//rwHZIrogkwaTAnJH+XOF5hSEX9YRmKJ/CFq6r6mQ+CZSGLw96f9Bg
kPMa1McyDLb3p/G+BKi8NrrCVVr6UknIfsb0Ibr/MlYqrZ3ld80n+HLF+cZ83DFA
Pgc5sJnoSejehNSdLXd/2T4CrKFnewercYKBxArnmsRjoXp3DcKdjsnsifSRuY0B
kSW6gs0M2IUj+jJQObZkVoUvOFSmKLKIlpSDlOKquAKj4bdUIRcKGTjryYYCnS9g
uPQJLs9urDA7sDktxA95zi9QxoB9oCwFTwg5DC6nZp3KIwnPJqlqd6Vy4zlfWZTh
8+xvDntN0nj3H/qHrA==
-----END CERTIFICATE-----
`
