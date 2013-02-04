// Copyright 2013 Andras Belicza. All rights reserved.

// Button component interface and implementation.

package gwu

// Button interface defines a clickable button.
// 
// Suggested event type to handle actions: ETYPE_CLICK
// 
// Default style class: "gwu-Button"
type Button interface {
	// Button is a component.
	Comp

	// Button has text.
	HasText

	// Button can be enabled/disabled.
	HasEnabled
}

// Button implementation.
type buttonImpl struct {
	compImpl       // Component implementation
	hasTextImpl    // Has text implementation
	hasEnabledImpl // Has enabled implementation
}

// NewButton creates a new Button.
func NewButton(text string) Button {
	c := newButtonImpl("", text)
	c.Style().AddClass("gwu-Button")
	return &c
}

// newButtonImpl creates a new buttonImpl.
func newButtonImpl(valueProviderJs string, text string) buttonImpl {
	return buttonImpl{newCompImpl(valueProviderJs), newHasTextImpl(text), newHasEnabledImpl()}
}

func (c *buttonImpl) Render(w writer) {
	w.Writes("<button type=\"button\"")
	c.renderAttrsAndStyle(w)
	c.renderEHandlers(w)
	c.renderEnabled(w)
	w.Write(_STR_GT)

	c.renderText(w)

	w.Writes("</button>")
}
