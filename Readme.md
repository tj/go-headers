# Header

Header format parser to match Netlify's [format](https://www.netlify.com/docs/headers-and-basic-auth/).

## Example

```go
/*
  Content-Type: text/plain
  X-Foo: 1
  X-bar: 2

# Some comment
/admin/*
  WWW-Authenticate: Basic whatever

/docs/*
  X-Foo: 1
X-bar   : 1
```

yields

```json
{
  "/*": {
    "Content-Type": [
      "text/plain"
    ],
    "X-Bar": [
      "2"
    ],
    "X-Foo": [
      "1"
    ]
  },
  "/admin/*": {
    "Www-Authenticate": [
      "Basic whatever"
    ]
  },
  "/docs/*": {
    "X-Bar": [
      "1"
    ],
    "X-Foo": [
      "1"
    ]
  }
}
```

---

[![GoDoc](https://godoc.org/github.com/tj/headers-go?status.svg)](https://godoc.org/github.com/tj/headers-go)
![](https://img.shields.io/badge/license-MIT-blue.svg)
![](https://img.shields.io/badge/status-stable-green.svg)

<a href="https://apex.sh"><img src="http://tjholowaychuk.com:6000/svg/sponsor"></a>
