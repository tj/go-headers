package headers

import (
	"bufio"
	"io"
	"net/http"
	"strings"

	"github.com/pkg/errors"
)

// Must parse utility.
func Must(m map[string]http.Header, err error) map[string]http.Header {
	if err != nil {
		panic(err)
	}

	return m
}

// Parse the given reader.
func Parse(r io.Reader) (map[string]http.Header, error) {
	rules := make(map[string]http.Header)
	var fields http.Header
	var path string

	s := bufio.NewScanner(r)

	for s.Scan() {
		line := strings.TrimSpace(s.Text())

		// empty
		if line == "" {
			continue
		}

		// comment
		if strings.HasPrefix(line, "#") {
			continue
		}

		// field
		if strings.Contains(line, ":") {
			if fields == nil {
				return nil, errors.New("path must precede fields")
			}

			parts := strings.Split(line, ":")
			k := strings.TrimSpace(parts[0])
			v := strings.TrimSpace(parts[1])
			fields.Add(k, v)
			continue
		}

		// path
		if path != "" && fields != nil {
			rules[path] = fields
		}

		path = line
		fields = make(http.Header)
	}

	if fields != nil {
		rules[path] = fields
	}

	return rules, s.Err()
}

// ParseString parses the given string.
func ParseString(s string) (map[string]http.Header, error) {
	return Parse(strings.NewReader(s))
}
