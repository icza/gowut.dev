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

import (
	"time"
)

// SessMonitor interface defines a component which monitors and displays
// the session timeout and network connectivity at client side without
// interacting with the session.
//
// Default style classes: "gwu-SessMonitor", "gwu-SessMonitor-Expired",
// ".gwu-SessMonitor-Error"
type SessMonitor interface {
	// SessMonitor is a component.
	Comp

	// Interval returns the duration between session checks.
	Interval() time.Duration

	// SetInterval sets the duration between session checks.
	//
	// Note: while this method allows you to pass an arbitrary time.Duration,
	// implementation might be using less precision (most likely millisecond).
	// Durations less than 1 ms might be rounded up to 1 ms.
	SetInterval(interval time.Duration)

	// Active tells if the session monitor is active.
	Active() bool

	// SetActive sets if the session monitor is active.
	// If the session monitor is not active, session checks will not be performed.
	// If a session monitor is deactivated and activated again, its interval is reset.
	SetActive(active bool)

	// SetJsConverter sets the Javascript function name which converts
	// a float second time value to a displayable string.
	// The default value is "convertSessTimeout" whose implementation is:
	//     function convertSessTimeout(sec) {
	//         if (sec <= 0)
	//             return "Expired!";
	//         else if (sec < 60)
	//             return "<1 min";
	//         else
	//             return "~" + Math.round(sec / 60) + " min";
	//     }
	SetJsConverter(jsFuncName string)

	// JsConverter returns the name of the Javascript function which converts
	// float second time values to displayable strings.
	JsConverter() string
}

// SessMonitor implementation
type sessMonitorImpl struct {
	compImpl // Component implementation

	interval time.Duration // Duration between session checks
	active   bool          // Tells if the timer is active
}

// NewSessMonitor creates a new SessMonitor.
func NewSessMonitor() SessMonitor {
	c := &sessMonitorImpl{newCompImpl(nil), time.Minute, true}
	c.Style().AddClass("gwu-SessMonitor")
	c.SetJsConverter("convertSessTimeout")
	return c
}

func (c *sessMonitorImpl) Interval() time.Duration {
	return c.interval
}

func (c *sessMonitorImpl) SetInterval(interval time.Duration) {
	if interval < time.Millisecond {
		interval = time.Millisecond
	}
	c.interval = interval
}

func (c *sessMonitorImpl) Active() bool {
	return c.active
}

func (c *sessMonitorImpl) SetActive(active bool) {
	c.active = active
}

func (c *sessMonitorImpl) SetJsConverter(jsFuncName string) {
	c.SetAttr("gwuJsFuncName", jsFuncName)
}

func (c *sessMonitorImpl) JsConverter() string {
	return c.Attr("gwuJsFuncName")
}

var (
	strEmptySpan     = []byte("<span></span>") // "<span></span>"
	strJsCheckSessOp = []byte("checkSession(") // "checkSession("
)

func (c *sessMonitorImpl) Render(w Writer) {
	w.Write(strSpanOp)
	c.renderAttrsAndStyle(w)
	c.renderEHandlers(w)
	w.Write(strGT)

	w.Write(strEmptySpan) // Placeholder for timeout value

	// <script>setupTimer(compId,"checkSession(compId)",interval,true,active,0);checkSession(compId);</script>
	w.Write(strScriptOp)
	w.Writev(int(c.id))
	w.Write(strComma)
	// js param
	w.Write(strQuote)
	w.Write(strJsCheckSessOp)
	w.Writev(int(c.id))
	w.Write(strJsParamCl)
	// end of js param
	w.Write(strComma)
	w.Writev(int(c.interval / time.Millisecond))
	w.Write(strComma)
	w.Writev(true) // Repeat
	w.Write(strComma)
	w.Writev(c.active) // Active
	w.Write(strComma)
	w.Writev(0) // Reset
	w.Write(strParenCl)
	w.Write(strSemicol)
	// Call sess check right away:
	w.Write(strJsCheckSessOp)
	w.Writev(int(c.id))
	w.Write(strScriptCl)

	w.Write(strSpanCl)
}
