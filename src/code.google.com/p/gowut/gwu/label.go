// Copyright 2013 Andras Belicza. All rights reserved.

// Label component interface and implementation.

package gwu

// Label interface defines a component which wraps a text into a component.
// 
// Default style class: "gwu-Label"
type Label interface {
	// Label is a component.
	Comp

	// Label has text.
	HasText
}

// Label implementation
type labelImpl struct {
	compImpl    // Component implementation
	hasTextImpl // Has text implementation
}

// NewLabel creates a new Label component.
func NewLabel(text string) Label {
	c := &labelImpl{newCompImpl(""), newHasTextImpl(text)}
	c.Style().AddClass("gwu-Label")
	return c
}

func (c *labelImpl) Render(w writer) {
	w.Write(_STR_SPAN_OP)
	c.renderAttrsAndStyle(w)
	c.renderEHandlers(w)
	w.Write(_STR_GT)

	c.renderText(w)

	w.Write(_STR_SPAN_CL)
}
