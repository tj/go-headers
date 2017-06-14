package headers_test

import (
	"encoding/json"
	"os"

	"github.com/tj/go-headers"
)

func Example() {
	h := headers.Must(headers.ParseString(`
    # Comment
    /*
      Content-Type: text/plain
      X-Foo: 1
      X-bar: 2

    # Moar comment
    # Comment
    /admin/*
      WWW-Authenticate: Basic whatever

    /docs/*
      X-Foo: 1
      X-Foo: 2
    X-bar   : 1
  `))

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	enc.Encode(h)
	// Output:
	// {
	//   "/*": {
	//     "Content-Type": [
	//       "text/plain"
	//     ],
	//     "X-Bar": [
	//       "2"
	//     ],
	//     "X-Foo": [
	//       "1"
	//     ]
	//   },
	//   "/admin/*": {
	//     "Www-Authenticate": [
	//       "Basic whatever"
	//     ]
	//   },
	//   "/docs/*": {
	//     "X-Bar": [
	//       "1"
	//     ],
	//     "X-Foo": [
	//       "1",
	//       "2"
	//     ]
	//   }
	// }
}
