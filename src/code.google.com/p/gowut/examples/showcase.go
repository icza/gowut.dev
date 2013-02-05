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

// A Gowut "Showcase of Features" application.

package main

import (
	"code.google.com/p/gowut/gwu"
	"fmt"
)

func plural(i int) string {
	if i == 1 {
		return ""
	}
	return "s"
}

func buildCheckBoxShow() gwu.Comp {
	p := gwu.NewPanel()

	suml := gwu.NewLabel("")

	p.Add(gwu.NewLabel("Check the days you want to work on:"))

	cbs := []gwu.CheckBox{gwu.NewCheckBox("Monday"), gwu.NewCheckBox("Tuesday"), gwu.NewCheckBox("Wednesday"),
		gwu.NewCheckBox("Thursday"), gwu.NewCheckBox("Friday"), gwu.NewCheckBox("Saturday"), gwu.NewCheckBox("Sunday")}
	cbs[5].SetEnabled(false)
	cbs[6].SetEnabled(false)

	for _, cb := range cbs {
		p.Add(cb)
		cb.AddEHandlerFunc(func(e gwu.Event) {
			sum := 0
			for _, cb2 := range cbs {
				if cb2.State() {
					sum++
				}
			}
			suml.SetText(fmt.Sprintf("%d day%s is a total of %d hours a week.", sum, plural(sum), sum*8))
			e.MarkDirty(suml)
		}, gwu.ETYPE_CLICK)
	}

	p.Add(suml)

	return p
}

func buildListBoxShow() gwu.Comp {
	p := gwu.NewPanel()
	p.Add(gwu.NewLabel("TODO"))
	return p
}

func buildTextBoxShow() gwu.Comp {
	p := gwu.NewPanel()
	p.Add(gwu.NewLabel("TODO"))
	return p
}

func buildPasswBoxShow() gwu.Comp {
	p := gwu.NewPanel()
	p.Add(gwu.NewLabel("TODO"))
	return p
}

func buildRadioButtonShow() gwu.Comp {
	p := gwu.NewPanel()

	p.Add(gwu.NewLabel("Select your favorite programming language:"))

	group := gwu.NewRadioGroup("lang")
	rbs := []gwu.RadioButton{gwu.NewRadioButton("Go", group), gwu.NewRadioButton("Java", group), gwu.NewRadioButton("C", group),
		gwu.NewRadioButton("C++", group), gwu.NewRadioButton("QBasic (nah this can't be your favorite)", group)}
	rbs[4].SetEnabled(false)

	for _, rb := range rbs {
		p.Add(rb)
	}

	space := gwu.NewLabel("")
	p.Add(space)
	p.CellFmt(space).Style().SetHeightPx(20)

	p.Add(gwu.NewLabel("Select your favorite computer game:"))

	group = gwu.NewRadioGroup("game")
	rbs = []gwu.RadioButton{gwu.NewRadioButton("StarCraft - Broodwar", group), gwu.NewRadioButton("StarCraft II", group),
		gwu.NewRadioButton("Other", group)}

	for _, rb := range rbs {
		p.Add(rb)
	}

	return p
}

func buildSwitchButtonShow() gwu.Comp {
	p := gwu.NewPanel()
	p.Add(gwu.NewLabel("TODO"))
	return p
}

func buildLinkShow() gwu.Comp {
	p := gwu.NewPanel()
	p.Add(gwu.NewLabel("TODO"))
	return p
}

func buildPanelShow() gwu.Comp {
	p := gwu.NewPanel()
	p.Add(gwu.NewLabel("TODO"))
	return p
}

func buildTableShow() gwu.Comp {
	p := gwu.NewPanel()
	p.Add(gwu.NewLabel("TODO"))
	return p
}

func buildTabPanelShow() gwu.Comp {
	p := gwu.NewPanel()
	p.Add(gwu.NewLabel("TODO"))
	return p
}

func buildWindowShow() gwu.Comp {
	p := gwu.NewPanel()
	p.Add(gwu.NewLabel("TODO"))
	return p
}

func buildButtonShow() gwu.Comp {
	p := gwu.NewPanel()
	p.Add(gwu.NewLabel("TODO"))
	return p
}

func buildHtmlShow() gwu.Comp {
	p := gwu.NewPanel()
	p.Add(gwu.NewLabel("TODO"))
	return p
}

func buildImageShow() gwu.Comp {
	p := gwu.NewPanel()
	p.Add(gwu.NewLabel("TODO"))
	return p
}

func buildLabelShow() gwu.Comp {
	p := gwu.NewPanel()
	p.Add(gwu.NewLabel("TODO"))
	return p
}

func buildLinkShow2() gwu.Comp {
	p := gwu.NewPanel()
	p.Add(gwu.NewLabel("TODO"))
	return p
}

func buildShowcase(sess gwu.Session) {
	win := gwu.NewWindow("show", "Showcase of Features - Gowut")
	win.Style().SetFullSize()

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
	logout := gwu.NewLink("Logout", "#")
	logout.SetTarget("")
	logout.AddEHandlerFunc(func(e gwu.Event) {
		e.RemoveSess()
		e.ReloadWin("login")
	}, gwu.ETYPE_CLICK)
	header.Add(logout)
	header.CellFmt(logout).SetHAlign(gwu.HA_RIGHT)
	header.CellFmt(logout).Style().SetWidthPx(100)
	win.Add(header)

	content := gwu.NewPanel()
	content.SetLayout(gwu.LAYOUT_HORIZONTAL)
	content.SetVAlign(gwu.VA_TOP)
	content.Style().SetFullSize()

	showWrapper := gwu.NewPanel()
	selectedShow := gwu.NewLabel("")
	selectedShow.Style().SetFontWeight(gwu.FONT_WEIGHT_BOLD).SetFontSize("110%")
	showWrapper.Add(selectedShow)

	links := gwu.NewPanel()

	// Lazily initialized, cached "Show" components
	showComps := make(map[string]gwu.Comp)
	var selectedLink gwu.Label
	addShowLink := func(show string, buildFunc func() gwu.Comp) {
		link := gwu.NewLabel(show)
		link.Style().SetFullWidth().SetCursor(gwu.CURSOR_POINTER).SetDisplay(gwu.DISPLAY_BLOCK).SetColor(gwu.CLR_BLUE)
		link.AddEHandlerFunc(func(e gwu.Event) {
			if selectedLink != nil {
				selectedLink.Style().SetBackground("")
				e.MarkDirty(selectedLink)
				showWrapper.Remove(showComps[selectedLink.Text()])
			}
			selectedLink = link
			selectedLink.Style().SetBackground("#aaffaa")
			selectedShow.SetText(show)
			showComp := showComps[show]
			if showComp == nil {
				showComp = buildFunc()
				showComps[show] = showComp
			}
			showWrapper.Add(showComp)
			e.MarkDirty(selectedLink, showWrapper)
		}, gwu.ETYPE_CLICK)
		links.Add(link)
	}

	links.Style().SetFullSize().Set("border-right", "2px solid #777777")
	l = gwu.NewLabel("Component Palette")
	l.Style().SetFontWeight(gwu.FONT_WEIGHT_BOLD).SetFontSize("120%")
	links.Add(l)
	l = gwu.NewLabel("Input components")
	l.Style().SetFontWeight(gwu.FONT_WEIGHT_BOLD)
	links.Add(l)
	addShowLink("CheckBox", buildCheckBoxShow)
	addShowLink("ListBox", buildListBoxShow)
	addShowLink("TextBox", buildTextBoxShow)
	addShowLink("PasswBox", buildPasswBoxShow)
	addShowLink("RadioButton", buildRadioButtonShow)
	addShowLink("SwitchButton", buildSwitchButtonShow)
	l = gwu.NewLabel("Containers")
	l.Style().SetFontWeight(gwu.FONT_WEIGHT_BOLD)
	links.Add(l)
	addShowLink("Link", buildLinkShow)
	addShowLink("Panel", buildPanelShow)
	addShowLink("Table", buildTableShow)
	addShowLink("TabPanel", buildTabPanelShow)
	addShowLink("Window", buildWindowShow)
	l = gwu.NewLabel("Other components")
	l.Style().SetFontWeight(gwu.FONT_WEIGHT_BOLD)
	links.Add(l)
	addShowLink("Button", buildButtonShow)
	addShowLink("Html", buildHtmlShow)
	addShowLink("Image", buildImageShow)
	addShowLink("Label", buildLabelShow)
	addShowLink("Link", buildLinkShow2)
	filler := gwu.NewLabel("")
	links.Add(filler)
	links.CellFmt(filler).Style().SetFullSize()
	content.Add(links)
	content.CellFmt(links).Style().SetWidthPx(200)
	content.Add(showWrapper)

	win.Add(content)
	win.CellFmt(content).Style().SetFullSize()

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
			e.Session().RemoveWin(win)
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

	// Start GUI server
	if err := server.Start("login"); err != nil {
		fmt.Println("Error: Cound not start GUI server:", err)
		return
	}
}
