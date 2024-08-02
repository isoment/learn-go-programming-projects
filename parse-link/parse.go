package parselink

import "io"

// Represents a link in an HTML document
type Link struct {
	Href string
	Text string
}

func Parse(r io.Reader) ([]Link, error) {
	return nil, nil
}
