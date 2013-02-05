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

// A Gowut showcase of features application.

package main

import (
	"code.google.com/p/gowut/gwu"
	"fmt"
)

func buildShowcase(sess gwu.Session) {
	win := gwu.NewWindow("show", "Showcase of Features - Gowut")
	win.Style().SetFullWidth()

	header := gwu.NewPanel()
	header.SetLayout(gwu.LAYOUT_HORIZONTAL)
	header.Style().SetFullWidth().Set("border-bottom", "2px solid #777777")
	l := gwu.NewLabel("Gowut - Showcase of Features")
	l.Style().SetFontWeight(gwu.FONT_WEIGHT_BOLD).SetFontSize("120%")
	header.Add(l)
	themespan := gwu.NewPanel()
	themespan.SetLayout(gwu.LAYOUT_HORIZONTAL)
	themespan.Add(gwu.NewLabel("Theme:"))
	themes := gwu.NewListBox([]string{"default", "debug"})
	themes.AddEHandlerFunc(func(e gwu.Event) {
		win.SetTheme(themes.SelectedValue())
		e.ReloadWin("show")
	}, gwu.ETYPE_CHANGE)
	themespan.Add(themes)
	header.Add(themespan)
	header.CellFmt(themespan).SetHAlign(gwu.HA_RIGHT)
	win.Add(header)

	content := gwu.NewPanel()
	content.SetLayout(gwu.LAYOUT_HORIZONTAL)
	content.Style().SetFullWidth()

	links := gwu.NewPanel()
	links.Style().SetWidthPx(200).Set("border-right", "2px solid #777777")
	l = gwu.NewLabel("Component Palette")
	l.Style().SetFontWeight(gwu.FONT_WEIGHT_BOLD).SetFontSize("120%")
	links.Add(l)
	l = gwu.NewLabel("Input components")
	l.Style().SetFontWeight(gwu.FONT_WEIGHT_BOLD)
	links.Add(l)
	links.Add(gwu.NewLink("CheckBox", "#"))
	links.Add(gwu.NewLink("Listbox", "#"))
	links.Add(gwu.NewLink("TextBox", "#"))
	links.Add(gwu.NewLink("PasswBox", "#"))
	links.Add(gwu.NewLink("RadioButton", "#"))
	links.Add(gwu.NewLink("SwitchButton", "#"))
	l = gwu.NewLabel("Containers")
	l.Style().SetFontWeight(gwu.FONT_WEIGHT_BOLD)
	links.Add(l)
	links.Add(gwu.NewLink("Link", "#"))
	links.Add(gwu.NewLink("Panel", "#"))
	links.Add(gwu.NewLink("Table", "#"))
	links.Add(gwu.NewLink("TabPanel", "#"))
	links.Add(gwu.NewLink("Window", "#"))
	l = gwu.NewLabel("Other components")
	l.Style().SetFontWeight(gwu.FONT_WEIGHT_BOLD)
	links.Add(l)
	links.Add(gwu.NewLink("Button", "#"))
	links.Add(gwu.NewLink("Html", "#"))
	links.Add(gwu.NewLink("Image", "#"))
	links.Add(gwu.NewLink("Label", "#"))
	links.Add(gwu.NewLink("Link", "#"))
	content.Add(links)

	win.Add(content)

	footer := gwu.NewPanel()
	footer.SetLayout(gwu.LAYOUT_HORIZONTAL)
	footer.Style().SetFullWidth().Set("border-top", "2px solid #777777")
	l = gwu.NewLabel("Gowut version " + gwu.GOWUT_VERSION)
	l.Style().SetFontStyle(gwu.FONT_STYLE_ITALIC)
	footer.Add(l)
	footer.CellFmt(l).SetHAlign(gwu.HA_RIGHT)
	win.Add(footer)

	sess.AddWin(win)
}

func buildLoginWin(sess gwu.Session) {
	win := gwu.NewWindow("login", "Login - Showcase of Features - Gowut")
	win.Style().SetFullWidth()
	win.SetHAlign(gwu.HA_CENTER)

	l := gwu.NewLabel("Gowut - Showcase of Features")
	l.Style().SetFontWeight(gwu.FONT_WEIGHT_BOLD).SetFontSize("150%")
	win.Add(l)
	l = gwu.NewLabel("Login")
	l.Style().SetFontWeight(gwu.FONT_WEIGHT_BOLD).SetFontSize("130%")
	win.Add(l)
	l = gwu.NewLabel("user/pass: admin/a")
	l.Style().SetFontSize("80%").SetFontStyle(gwu.FONT_STYLE_ITALIC)
	win.Add(l)

	errL := gwu.NewLabel("")
	errL.Style().SetColor(gwu.CLR_RED)
	win.Add(errL)

	table := gwu.NewTable()
	table.EnsureSize(2, 2)
	table.Add(gwu.NewLabel("User name:"), 0, 0)
	tb := gwu.NewTextBox("")
	table.Add(tb, 0, 1)
	table.Add(gwu.NewLabel("Password:"), 1, 0)
	pb := gwu.NewPasswBox("")
	table.Add(pb, 1, 1)
	win.Add(table)
	b := gwu.NewButton("OK")
	b.AddEHandlerFunc(func(e gwu.Event) {
		if tb.Text() == "admin" && pb.Text() == "a" {
			e.Session().RemoveWin(win) // Login win is removed, password will not be retrievable from the browser
			buildShowcase(e.Session())
			e.ReloadWin("show")
		} else {
			e.SetFocusedComp(tb)
			errL.SetText("Invalid user name or password!")
			e.MarkDirty(errL)
		}
	}, gwu.ETYPE_CLICK)
	win.Add(b)

	win.SetFocusedCompId(tb.Id())

	sess.AddWin(win)
}

type SessHandler struct{}

func (h SessHandler) Created(s gwu.Session) {
	buildLoginWin(s)
}

func (h SessHandler) Removed(s gwu.Session) {}

func main() {
	// Create GUI server
	server := gwu.NewServer("showcase", "")
	server.SetText("Gowut - Showcase of Features")

	server.AddSessCreatorName("login", "Login Window")
	server.AddSHandler(SessHandler{})

	//server.SetTheme("debug")

	// Start GUI server
	if err := server.Start("login"); err != nil {
		fmt.Println("Error: Cound not start GUI server:", err)
		return
	}
}
