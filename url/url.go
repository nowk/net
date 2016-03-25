package url

import (
	"net/url"
	"regexp"
	"strings"
)

type URL struct {
	*url.URL

	// Port is the :8080 without the :
	Port string

	// HostOnly is the host only, this will never include any port information
	HostOnly string
}

var (
	schemePrefix = regexp.MustCompile("^([a-z]+:)?(//)")

	// shortcut to url.Parse
	parse = url.Parse
)

func Parse(str string) (*URL, error) {
	// force anonymous porotocol to allow url.Parse to parse the string
	// properly
	if !schemePrefix.MatchString(str) {
		str = "//" + str
	}

	// parse the url, through net/url
	neturl, err := parse(str)
	if err != nil {
		return nil, err
	}

	var (
		host, port = parseHostPort(neturl.Host)

		u = &URL{
			URL: neturl,

			Port:     port,
			HostOnly: host,
		}
	)

	return u, nil
}

// parseHostPort parses the host:port at :. This is similar to net.SplitHostPort
// but does not error when there is no port found, port is just empty.
func parseHostPort(str string) (string, string) {
	var (
		host string
		port string

		i = strings.Index(str, ":")
	)
	if i == -1 {
		return str, ""
	}

	host = str[:i]
	port = str[i+1:]

	return host, port
}
