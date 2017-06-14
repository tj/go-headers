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

		/*
		  Link: </style.css>; rel=preload; as=stylesheet
		  Link: </main.js>; rel=preload; as=script
		  Link: </image.jpg>; rel=preload; as=image
  `))

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	enc.Encode(h)
	// Output:
	// {
	//   "/*": {
	//     "Link": [
	//       "\u003c/style.css\u003e; rel=preload; as=stylesheet",
	//       "\u003c/main.js\u003e; rel=preload; as=script",
	//       "\u003c/image.jpg\u003e; rel=preload; as=image"
	//     ],
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
