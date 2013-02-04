// Copyright 2013 Andras Belicza. All rights reserved.

// Image component interface and implementation.

package gwu

// Image interface defines an image.
// 
// Default style class: "gwu-Image"
type Image interface {
	// Image is a component.
	Comp

	// Image has text which is its description (alternate text).
	HasText

	// Image has URL string.
	HasUrl
}

// Image implementation
type imageImpl struct {
	compImpl    // Component implementation
	hasTextImpl // Has text implementation
	hasUrlImpl  // Has text implementation
}

// NewImage creates a new Image component.
func NewImage(text, url string) Image {
	c := &imageImpl{newCompImpl(""), newHasTextImpl(text), newHasUrlImpl(url)}
	c.Style().AddClass("gwu-Image")
	return c
}

func (c *imageImpl) Render(w writer) {
	w.Writes("<img")
	c.renderUrl("src", w)
	c.renderAttrsAndStyle(w)
	c.renderEHandlers(w)
	w.Writes(" alt=\"")
	c.renderText(w)
	w.Writes("\">")
}
