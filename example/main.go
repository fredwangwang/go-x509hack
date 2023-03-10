package main

import (
	"crypto/x509"
	"encoding/pem"
	"log"
	// _ "github.com/fredwangwang/go-x509hack" // <- uncomment this line
)

var cert = `-----BEGIN CERTIFICATE-----
MIIELjCCAxagAwIBAgIFIABlIQQwDQYJKoZIhvcNAQEFBQAwWTELMAkGA1UEBhMC
Q04xMDAuBgNVBAoTJ0NoaW5hIEZpbmFuY2lhbCBDZXJ0aWZpY2F0aW9uIEF1dGhv
cml0eTEYMBYGA1UEAxMPQ0ZDQSBURVNUIE9DQTExMB4XDTEzMTIxODA5MjcwN1oX
DTE1MTIxODA5MjcwN1owazELMAkGA1UEBhMCQ04xDjAMBgNVBAoTBXl1emhpMRIw
EAYDVQQLEwlDTUJDX0RDTVMxFDASBgNVBAsTC0luZGl2aWR1YWwxMSIwIAYDVQQD
FBljbj0wMDA0MDAwMDpTSUdOQDAwMDAwMDAyMIIBIjANBgkqhkiG9w0BAQEFAAOC
AQ8AMIIBCgKCAQEAtTBhimpRdnIMC7SBfl3xeSJR8XF0Jqf9XItroLMoffxyNfRB
khKGf87oLDaGfZa6s/+vct2iktHNMcjS8gMmlBj/RjXflAdEVkeQAfSiZFN8Bn9q
14+dEXtVbvrosaFHmfTmHWchMMeXMJlg9K6akM+lVwmPjLfT4bcDk67VVUCy4uLj
xe0YgrBFfqAxcj4/vlHAdKY+3n+keUQ9PrBRuYXQcGgXaBEGlhnCJIrnVnnZdRUS
T3qt6iUUUu7/2rIvvrv+/aLN3li432LuhJc3DWANhwucC7DDpXoFScoqrgpQbofc
vPdbJdsRrLvZ/XYDmmcdYOMsQ7AnZIbVVUws3wIDAQABo4HqMIHnMB8GA1UdIwQY
MBaAFPwLvESaDjGhg6mBhyceBULGv1b4MEgGA1UdIARBMD8wPQYIYIEchu8qAQIw
MTAvBggrBgEFBQcCARYjaHR0cDovL3d3dy5jZmNhLmNvbS5jbi91cy91cy0xNS5o
dG0wOQYDVR0fBDIwMDAuoCygKoYoaHR0cDovLzIxMC43NC40MS44Ny9vY2ExMS9S
U0EvY3JsNjM5LmNybDALBgNVHQ8EBAMCA+gwHQYDVR0OBBYEFLBo5Ut5b4js+bsL
yVe73j4xNcDOMBMGA1UdJQQMMAoGCCsGAQUFBwMCMA0GCSqGSIb3DQEBBQUAA4IB
AQBQbuGIjDNd0wT1dGcjsImtXAmHZaytqUn3mY2VdAEZhR1rLEGjkG+j30zIV7K8
BrUWOf0Pny6tQ+OTevNGCV1hiO2ZQryB50R/eBGSk70dBqIDzpOk4cAj6OBaO+uh
Uqw2jdDOt1rV3UeptMD4Mbs3FhdQy3sDApkZs10px/5PHMohjktnwSuFmXaNwYsX
Sv282c6tjiDWYcQrTHZVfDhhrooAjLSQ6wDWW/QnjhbcNbY2lnf9hReYoDBloK8G
O0AH1nfVkMVY0Ja6OdwnJD+NFIL0yPqz1gH7qzrmcGOZefiUPC+LGU0D+qYU1f22
afnVaYT41MCjuM1nXmMcEmLt
-----END CERTIFICATE-----`

func main() {
	block, _ := pem.Decode([]byte(cert))
	certBytes := block.Bytes
	cert, err := x509.ParseCertificate(certBytes)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(cert.Subject.OrganizationalUnit)
}
