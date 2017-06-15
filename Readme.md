# Header

Header format parser to match Netlify's [format](https://www.netlify.com/docs/headers-and-basic-auth/).

## Example

```go
/*
  X-Frame-Options: DENY
  X-XSS-Protection: 1; mode=block

## A path:
/templates/index.html
  # Headers for that path:
  X-Frame-Options: DENY
  X-XSS-Protection: 1; mode=block

/templates/index2.html
  X-Frame-Options: SAMEORIGIN
```

yields

```json
{
  "/*": {
    "X-Frame-Options": [
      "DENY"
    ],
    "X-Xss-Protection": [
      "1; mode=block"
    ]
  },
  "/templates/index.html": {
    "X-Frame-Options": [
      "DENY"
    ],
    "X-Xss-Protection": [
      "1; mode=block"
    ]
  },
  "/templates/index2.html": {
    "X-Frame-Options": [
      "SAMEORIGIN"
    ]
  }
}
```

---

[![GoDoc](https://godoc.org/github.com/tj/go-headers?status.svg)](https://godoc.org/github.com/tj/go-headers)
![](https://img.shields.io/badge/license-MIT-blue.svg)
![](https://img.shields.io/badge/status-stable-green.svg)

<a href="https://apex.sh"><img src="http://tjholowaychuk.com:6000/svg/sponsor"></a>
