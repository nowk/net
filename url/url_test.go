package url

import (
	"net/url"
	"reflect"
	"testing"
)

func TestParseReturnsAUrlWithPort(t *testing.T) {
	var cases = [][]string{
		{"http://example.com:8080", "http"},
		{"HTTP://example.com:8080", "http"},
		{"//example.com:8080", ""},
		{"example.com:8080", ""},
	}

	for _, v := range cases {
		var (
			giv = v[0]

			u, err = Parse(giv)
		)
		if err != nil {
			t.Fatal(err)
		}

		var (
			exp = &URL{
				URL: &url.URL{
					Scheme: v[1],
					Host:   "example.com:8080",
				},
				Port:     "8080",
				HostOnly: "example.com",
			}

			got = u
		)
		if !reflect.DeepEqual(exp, got) {
			t.Errorf("expected %s, got %s", exp, got)
		}
	}
}

func TestParseReturnsAUrl(t *testing.T) {
	var cases = [][]string{
		{"http://example.com", "http"},
		{"HTTP://example.com", "http"},
		{"//example.com", ""},
		{"example.com", ""},
	}

	for _, v := range cases {
		var (
			giv = v[0]

			u, err = Parse(giv)
		)
		if err != nil {
			t.Fatal(err)
		}

		var (
			exp = &URL{
				URL: &url.URL{
					Scheme: v[1],
					Host:   "example.com",
				},
				Port:     "",
				HostOnly: "example.com",
			}

			got = u
		)
		if !reflect.DeepEqual(exp, got) {
			t.Errorf("expected %s, got %s", exp, got)
		}
	}
}
