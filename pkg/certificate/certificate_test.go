package certificate

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestParse(t *testing.T) {
	testCases := []struct {
		name     string
		cert     string
		expected string
	}{
		{
			name: "case 0: parse a typical certificate, without newlines",
			cert: "-----BEGIN CERTIFICATE----- something otherthing lastthing -----END CERTIFICATE-----",
			expected: `-----BEGIN CERTIFICATE-----
something
otherthing
lastthing
-----END CERTIFICATE-----`,
		},
		{
			name: "case 0: parse a typical certificate, with newlines",
			cert: `-----BEGIN CERTIFICATE-----
something
otherthing
lastthing
-----END CERTIFICATE-----`,
			expected: `-----BEGIN CERTIFICATE-----
something
otherthing
lastthing
-----END CERTIFICATE-----`,
		},
		{
			name: "case 2: parse a typical certificate, encoded in yaml",
			cert: "        | -----BEGIN CERTIFICATE----- something otherthing lastthing -----END CERTIFICATE-----",
			expected: `-----BEGIN CERTIFICATE-----
something
otherthing
lastthing
-----END CERTIFICATE-----`,
		},
		{
			name: "case 3: parse a yaml-encoded string",
			cert: "        | something otherthing lastthing",
			expected: `-----BEGIN CERTIFICATE-----
something
otherthing
lastthing
-----END CERTIFICATE-----`,
		},
		{
			name: "case 4: parse a regular string",
			cert: "something otherthing lastthing",
			expected: `-----BEGIN CERTIFICATE-----
something
otherthing
lastthing
-----END CERTIFICATE-----`,
		},
		{
			name: "case 5: parse an empty string",
			cert: "",
			expected: `-----BEGIN CERTIFICATE-----

-----END CERTIFICATE-----`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := Parse(tc.cert)
			diff := cmp.Diff(result, tc.expected)
			if len(diff) > 0 {
				t.Fatalf("value not expected, got:\n %s", diff)
			}
		})
	}
}
