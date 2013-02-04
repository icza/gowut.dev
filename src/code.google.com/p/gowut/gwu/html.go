// Copyright 2013 Andras Belicza. All rights reserved.

// Defines the Html component.

package gwu

// Html interface defines a component which wraps an HTML text into a component.
// 
// Default style class: "gwu-Html"
type Html interface {
	// Html is a component.
	Comp

	// Html returns the HTML text.
	Html() string

	// SetHtml sets the HTML text.
	SetHtml(html string)
}

// Html implementation
type htmlImpl struct {
	compImpl // Component implementation

	html string // HTML text
}

// NewHtml creates a new Html component.
func NewHtml(html string) Html {
	c := &htmlImpl{newCompImpl(""), html}
	c.Style().AddClass("gwu-Html")
	return c
}

func (c *htmlImpl) Html() string {
	return c.html
}

func (c *htmlImpl) SetHtml(html string) {
	c.html = html
}

func (c *htmlImpl) Render(w writer) {
	w.Write(_STR_SPAN_OP)
	c.renderAttrsAndStyle(w)
	c.renderEHandlers(w)
	w.Write(_STR_GT)

	w.Writes(c.html)

	w.Write(_STR_SPAN_CL)
}
