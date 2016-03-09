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

// Session monitor component interface and implementation.

package gwu

// SessMonitor interface defines a component which monitors and displays
// the timeout of the session at client side.
//
// Default style classes: "gwu-SessMonitor", "gwu-SessMonitor-Expired"
type SessMonitor interface {
	// SessMonitor is a component.
	Comp
}

// Label implementation
type sessMonitorImpl struct {
	compImpl // Component implementation
}

// NewSessMonitor creates a new SessMonitor.
func NewSessMonitor() SessMonitor {
	c := &sessMonitorImpl{newCompImpl(nil)}
	c.Style().AddClass("gwu-SessMonitor")
	return c
}

var (
	strEmptySpan     = []byte("<span></span>")  // "<span></span>"
	strJsCheckSessOp = []byte(`"checkSession(`) // `"se(null,`
)

func (c *sessMonitorImpl) Render(w Writer) {
	w.Write(strSpanOp)
	c.renderAttrsAndStyle(w)
	c.renderEHandlers(w)
	w.Write(strGT)

	w.Write(strEmptySpan) // Placeholder for timeout value

	w.Write(strScriptOp)
	w.Writev(int(c.id))
	w.Write(strComma)
	// js param
	w.Write(strJsCheckSessOp)
	w.Writev(int(c.id))
	w.Write(strJsParamCl)
	// end of js param
	w.Write(strComma)
	w.Writev(60 * 1000) // 1 min
	w.Write(strComma)
	w.Writev(true) // Repeat
	w.Write(strComma)
	w.Writev(true) // Active
	w.Write(strComma)
	w.Writev(false) // Reset
	w.Write(strScriptCl)

	// Call sess check right away:
	// TODO
	w.Writevs("<script>checkSession(", int(c.id), ");</script>")

	w.Write(strSpanCl)

}
