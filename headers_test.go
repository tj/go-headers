package headers_test

import (
	"encoding/json"
	"os"

	"github.com/tj/go-headers"
)

func Example() {
	h := headers.Must(headers.ParseString(`
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
  `))

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	enc.Encode(h)
	// Output:
	// {
	//   "/*": {
	//     "X-Frame-Options": [
	//       "DENY"
	//     ],
	//     "X-Xss-Protection": [
	//       "1; mode=block"
	//     ]
	//   },
	//   "/templates/index.html": {
	//     "X-Frame-Options": [
	//       "DENY"
	//     ],
	//     "X-Xss-Protection": [
	//       "1; mode=block"
	//     ]
	//   },
	//   "/templates/index2.html": {
	//     "X-Frame-Options": [
	//       "SAMEORIGIN"
	//     ]
	//   }
	// }
}
