# net/url

[![Build Status](https://travis-ci.org/nowk/net.svg?branch=master)][0]
[![GoDoc](https://godoc.org/gopkg.in/nowk/net.v0?status.svg)][1]

  [0]: https://travis-ci.org/nowk/net
  [1]: http://godoc.org/gopkg.in/nowk/net.v0

net/url provides a `Parse` method which wraps the original `net/url`'s own `Parse` to properly parse `host:port` when there is no protocol signature.

## Install

    go get github.com/nowk/net/url

## Usage

    import (
        "github.com/nowk/net/url"
    )

    u, err := url.Parse("example.com:8080")
    if err != nil {
        // handle
    }

    // &url.URL{
    //   URL: &net.URL{
    //     Scheme: "",
    //     Host: "example.com:8080",
    //     ...
    //   },
    //   HostOnly: "example.com",
    //   Port:     "8080",
    // }


### License

MIT
