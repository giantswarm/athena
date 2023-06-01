package certificate

import (
	"fmt"
	"strings"
)

const (
	yamlMultiLinePrefix = "        | "

	certPrefix = "-----BEGIN CERTIFICATE-----"
	certSuffix = "-----END CERTIFICATE-----"
)

// Parse understands and encodes a PEM certificate into a string.
func Parse(cert string) string {
	cert = strings.TrimPrefix(cert, yamlMultiLinePrefix)

	cert = strings.TrimSpace(cert)
	cert = strings.TrimPrefix(cert, certPrefix)
	cert = strings.TrimSuffix(cert, certSuffix)
	cert = strings.TrimSpace(cert)

	cert = strings.ReplaceAll(cert, " ", "\n")

	cert = fmt.Sprintf("%s\n%s\n%s", certPrefix, cert, certSuffix)

	return cert
}
