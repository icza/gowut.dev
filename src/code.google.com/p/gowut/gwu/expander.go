// Copyright (C) 2013 Andras Belicza. All rights reserved.
// 
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
// 
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
// 
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

// Expander component interface and implementation.

package gwu

// Expander interface defines a component which can show or hide
// another component when clicked on the header.
// 
// Default style classes: "gwu-Expander", "gwu-Expander-Header",
// "gwu-Expander-Header-Open", "gwu-Expander-Content"
type Expander interface {
	// Expander is a Container.
	Container

	// Header returns the header component of the expander.
	Header() Comp

	// SetHeader sets the header component of the expander.
	SetHeader(h Comp)

	// Content returns the content component of the expander.
	Content() Comp

	// SetContent sets the content component of the expander.
	SetContent(c Comp)

	// Expanded returns whether the expander is expanded.
	Expanded() bool

	// SetExpanded sets whether the expander is expanded.
	SetExpanded(expanded bool)
}

// Expander implementation.
// Implemented
type expanderImpl struct {
	compImpl // Component implementation

	header   Comp // Header component
	content  Comp // Content component
	expanded bool // Tells whether the expander is expanded
}

// NewExpander creates a new Expander component.
// By default expanders are not expanded.
func NewExpander() Expander {
	c := &expanderImpl{compImpl: newCompImpl("")}
	c.Style().AddClass("gwu-Expander")
	return c
}

func (c *expanderImpl) Remove(c2 Comp) bool {
	if c.header.Equals(c2) {
		c2.setParent(nil)
		c.header = nil
		return true
	}

	if c.content.Equals(c2) {
		c2.setParent(nil)
		c.content = nil
		return true
	}

	return false
}

func (c *expanderImpl) ById(id ID) Comp {
	if c.id == id {
		return c
	}

	if c.header != nil {
		if c.header.Id() == id {
			return c.header
		}
		if c2, isContainer := c.header.(Container); isContainer {
			if c3 := c2.ById(id); c3 != nil {
				return c3
			}
		}
	}

	if c.content != nil {
		if c.content.Id() == id {
			return c.content
		}
		if c2, isContainer := c.content.(Container); isContainer {
			if c3 := c2.ById(id); c3 != nil {
				return c3
			}
		}
	}

	return nil
}

func (c *expanderImpl) Clear() {
	if c.header != nil {
		c.header.setParent(nil)
		c.header = nil
	}
	if c.content != nil {
		c.content.setParent(nil)
		c.content = nil
	}
}

func (c *expanderImpl) Header() Comp {
	return c.header
}

func (c *expanderImpl) SetHeader(header Comp) {
	c.header = header
}

func (c *expanderImpl) Content() Comp {
	return c.header
}

func (c *expanderImpl) SetContent(content Comp) {
	c.content = content
}

func (c *expanderImpl) Expanded() bool {
	return c.expanded
}

func (c *expanderImpl) SetExpanded(expanded bool) {
	c.expanded = expanded
}

func (c *expanderImpl) Render(w writer) {
	// TODO
	w.Writes("<a")
	c.renderAttrsAndStyle(w)
	c.renderEHandlers(w)
	w.Write(_STR_GT)

	if c.header != nil {
		c.header.Render(w)
	}

	w.Writes("</a>")
}
