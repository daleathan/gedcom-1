package html

import (
	"io"
	"strconv"
)

const (
	noBadgeCount = -1
)

// Card is a simple box with a header and body section.
type Card struct {
	title Component
	body  Component
	count int
}

func NewCard(title Component, count int, body Component) *Card {
	return &Card{
		title: title,
		body:  body,
		count: count,
	}
}

func (c *Card) WriteTo(w io.Writer) (int64, error) {
	var count = c.title
	if c.count != noBadgeCount {
		count = NewComponents(
			c.title,
			NewBadgePill("secondary", "float-right",
				NewText(strconv.Itoa(c.count))),
		)
	}

	heading := NewHeading(5, "card-header", count)
	components := NewComponents(heading, c.body)

	return NewDiv("card", components).WriteTo(w)
}
