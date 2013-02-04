// Copyright 2013 Andras Belicza. All rights reserved.

// Link component interface and implementation.

package gwu

// Link interface defines a clickable link pointing to a URL.
// Links are usually used with a text, although Link is a
// container, and allows to set a child component
// which if set will also be a part of the clickable link.
// 
// Default style class: "gwu-Link"
type Link interface {
	// Link is a Container.
	Container

	// Link has text.
	HasText

	// Link has URL string.
	HasUrl

	// Target returns the target of the link.
	Target() string

	// SetTarget sets the target of the link.
	// Tip: pass "_blank" if you want the URL to open in a new window.
	SetTarget(target string)

	// Comp returns the optional child component, if set.
	Comp() Comp

	// SetComp sets the only child component
	// (which can be a Container of course).
	SetComp(c Comp)
}

// Link implementation
type linkImpl struct {
	compImpl    // Component implementation
	hasTextImpl // Has text implementation
	hasUrlImpl  // Has text implementation

	comp Comp // Optional child component
}

// NewLink creates a new Link component.
// By default links open in a new window (tab)
// because their target is set to "_blank".
func NewLink(text, url string) Link {
	c := &linkImpl{newCompImpl(""), newHasTextImpl(text), newHasUrlImpl(url), nil}
	c.SetTarget("_blank")
	c.Style().AddClass("gwu-Link")
	return c
}

func (c *linkImpl) Remove(c2 Comp) bool {
	if c.comp == nil || !c.comp.Equals(c2) {
		return false
	}

	c2.setParent(nil)
	c.comp = nil

	return true
}

func (c *linkImpl) ById(id ID) Comp {
	if c.id == id {
		return c
	}

	if c.comp != nil {
		if c.comp.Id() == id {
			return c.comp
		}
		if c2, isContainer := c.comp.(Container); isContainer {
			if c3 := c2.ById(id); c3 != nil {
				return c3
			}
		}

	}

	return nil
}

func (c *linkImpl) Clear() {
	c.comp.setParent(nil)
	c.comp = nil
}

func (c *linkImpl) Target() string {
	return c.attrs["target"]
}

func (c *linkImpl) SetTarget(target string) {
	c.attrs["target"] = target
}

func (c *linkImpl) Comp() Comp {
	return c.comp
}

func (c *linkImpl) SetComp(c2 Comp) {
	c.comp = c2
}

func (c *linkImpl) Render(w writer) {
	w.Writes("<a")
	c.renderUrl("href", w)
	c.renderAttrsAndStyle(w)
	c.renderEHandlers(w)
	w.Write(_STR_GT)

	c.renderText(w)

	if c.comp != nil {
		c.comp.Render(w)
	}

	w.Writes("</a>")
}
