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

// plural returns an empty string if i is equal to 1,
// "s" otherwise.
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

	row := gwu.NewHorizontalPanel()
	l := gwu.NewLabel("Select a background color:")
	row.Add(l)
	lb := gwu.NewListBox([]string{"", "Black", "Red", "Green", "Blue", "White"})
	lb.AddEHandlerFunc(func(e gwu.Event) {
		l.Style().SetBackground(lb.SelectedValue())
		e.MarkDirty(l)
	}, gwu.ETYPE_CHANGE)
	row.Add(lb)
	p.Add(row)

	p.AddVSpace(10)
	p.Add(gwu.NewLabel("Select numbers that add up to 89:"))
	sumLabel := gwu.NewLabel("")
	lb2 := gwu.NewListBox([]string{"1", "2", "4", "8", "16", "32", "64", "128"})
	lb2.SetMulti(true)
	lb2.SetRows(10)
	lb2.AddEHandlerFunc(func(e gwu.Event) {
		sum := 0
		for _, idx := range lb2.SelectedIndices() {
			sum += 1 << uint(idx)
		}
		if sum == 89 {
			sumLabel.SetText("Hooray! You did it!")
		} else {
			sumLabel.SetText(fmt.Sprintf("Now quite there... (sum = %d)", sum))
		}
		e.MarkDirty(sumLabel)
	}, gwu.ETYPE_CHANGE)
	p.Add(lb2)
	p.Add(sumLabel)

	return p
}

func buildTextBoxShow() gwu.Comp {
	p := gwu.NewPanel()

	p.Add(gwu.NewLabel("Enter your name (max 15 characters):"))
	row := gwu.NewHorizontalPanel()
	tb := gwu.NewTextBox("")
	tb.SetMaxLength(15)
	tb.AddSyncOnETypes(gwu.ETYPE_KEY_UP)
	length := gwu.NewLabel("")
	length.Style().SetFontSize("80%").SetFontStyle(gwu.FONT_STYLE_ITALIC)
	tb.AddEHandlerFunc(func(e gwu.Event) {
		rem := 15 - len(tb.Text())
		length.SetText(fmt.Sprintf("(%d character%s left...)", rem, plural(rem)))
		e.MarkDirty(length)
	}, gwu.ETYPE_CHANGE, gwu.ETYPE_KEY_UP)
	row.Add(tb)
	row.Add(length)
	p.Add(row)

	p.AddVSpace(10)
	p.Add(gwu.NewLabel("Short biography:"))
	bio := gwu.NewTextBox("")
	bio.SetRows(5)
	bio.SetCols(40)
	p.Add(bio)

	p.AddVSpace(10)
	rtb := gwu.NewTextBox("This is just a read-only text box...")
	rtb.SetReadOnly(true)
	p.Add(rtb)

	p.AddVSpace(10)
	dtb := gwu.NewTextBox("...and a disabled one.")
	dtb.SetEnabled(false)
	p.Add(dtb)

	return p
}

func buildPasswBoxShow() gwu.Comp {
	p := gwu.NewPanel()

	p.Add(gwu.NewLabel("Enter your password:"))
	p.Add(gwu.NewPasswBox(""))

	return p
}

func buildRadioButtonShow() gwu.Comp {
	p := gwu.NewPanel()

	p.Add(gwu.NewLabel("Select your favorite programming language:"))

	group := gwu.NewRadioGroup("lang")
	rbs := []gwu.RadioButton{gwu.NewRadioButton("Go", group), gwu.NewRadioButton("Java", group), gwu.NewRadioButton("C / C++", group),
		gwu.NewRadioButton("Python", group), gwu.NewRadioButton("QBasic (nah this can't be your favorite)", group)}
	rbs[4].SetEnabled(false)

	for _, rb := range rbs {
		p.Add(rb)
	}

	p.AddVSpace(20)
	p.Add(gwu.NewLabel("Select your favorite computer game:"))

	group = gwu.NewRadioGroup("game")
	rbs = []gwu.RadioButton{gwu.NewRadioButton("StarCraft II", group), gwu.NewRadioButton("Minecraft", group),
		gwu.NewRadioButton("Other", group)}

	for _, rb := range rbs {
		p.Add(rb)
	}

	return p
}

func buildSwitchButtonShow() gwu.Comp {
	p := gwu.NewPanel()

	row := gwu.NewHorizontalPanel()
	row.Add(gwu.NewLabel("Here's an ON/OFF switch which enables/disables the other one:"))
	sw := gwu.NewSwitchButton()
	sw.SetOnOff("ENB", "DISB")
	sw.SetState(true)
	row.Add(sw)
	p.Add(row)

	row = gwu.NewHorizontalPanel()
	row.Add(gwu.NewLabel("And the other one:"))
	sw2 := gwu.NewSwitchButton()
	sw2.SetEnabled(true)
	sw2.Style().SetWidthPx(100)
	row.Add(sw2)
	sw.AddEHandlerFunc(func(e gwu.Event) {
		sw2.SetEnabled(sw.State())
		e.MarkDirty(sw2)
	}, gwu.ETYPE_CLICK)
	p.Add(row)

	return p
}

func buildLinkShow() gwu.Comp {
	p := gwu.NewPanel()

	link := gwu.NewLink("Obvious link to Google Home", "https://google.com")
	inside := gwu.NewPanel()
	inside.Style().SetBorder2(1, gwu.BRD_STYLE_SOLID, gwu.CLR_GRAY)
	inside.Add(gwu.NewLabel("Everything inside this box also links to Google!"))
	inside.Add(gwu.NewButton("Me too!"))
	link.SetComp(inside)
	p.Add(link)

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

	l := gwu.NewLabel("")

	btnp := gwu.NewHorizontalPanel()
	b := gwu.NewButton("Normal Button")
	b.AddEHandlerFunc(func(e gwu.Event) {
		switch e.Type() {
		case gwu.ETYPE_MOUSE_OVER:
			l.SetText("Mouse is over...")
		case gwu.ETYPE_MOUSE_OUT:
			l.SetText("Mouse is out.")
		case gwu.ETYPE_CLICK:
			x, y := e.Mouse()
			l.SetText(fmt.Sprintf("Clicked at x=%d, y=%d", x, y))
		}
		e.MarkDirty(l)
	}, gwu.ETYPE_CLICK, gwu.ETYPE_MOUSE_OVER, gwu.ETYPE_MOUSE_OUT)
	btnp.Add(b)

	b = gwu.NewButton("Disabled Button")
	b.SetEnabled(false)
	btnp.Add(b)

	p.Add(btnp)

	p.Add(l)

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

	p.Add(gwu.NewLink("Visit Gowut Home page", "https://sites.google.com/site/gowebuitoolkit/"))
	p.Add(gwu.NewLink("Visit Gowut Project page", "http://code.google.com/p/gowut/"))

	row := gwu.NewHorizontalPanel()
	row.Add(gwu.NewLabel("Discussion forum:"))
	row.Add(gwu.NewLink("https://groups.google.com/d/forum/gowebuitoolkit", "https://groups.google.com/d/forum/gowebuitoolkit"))
	p.Add(row)

	row = gwu.NewHorizontalPanel()
	row.Add(gwu.NewLabel("Send e-mail to the author of Gowut:"))
	email := "iczaaa" + "@" + "gmail.com"
	row.Add(gwu.NewLink("András Belicza <"+email+">", "mailto:"+email))
	p.Add(row)

	return p
}

func buildShowcase(sess gwu.Session) {
	win := gwu.NewWindow("show", "Showcase of Features - Gowut")
	win.Style().SetFullSize()

	header := gwu.NewHorizontalPanel()
	header.Style().SetFullWidth().SetBorderBottom2(2, gwu.BRD_STYLE_SOLID, "#777777").SetWhiteSpace(gwu.WHITE_SPACE_NOWRAP)
	l := gwu.NewLabel("Gowut - Showcase of Features")
	l.Style().SetFontWeight(gwu.FONT_WEIGHT_BOLD).SetFontSize("120%")
	header.Add(l)
	header.AddHConsumer()
	header.Add(gwu.NewLabel("Theme:"))
	themes := gwu.NewListBox([]string{"default", "debug"})
	themes.AddEHandlerFunc(func(e gwu.Event) {
		win.SetTheme(themes.SelectedValue())
		e.ReloadWin("show")
	}, gwu.ETYPE_CHANGE)
	header.Add(themes)
	header.AddHSpace(10)
	logout := gwu.NewLink("Logout", "#")
	logout.SetTarget("")
	logout.AddEHandlerFunc(func(e gwu.Event) {
		e.RemoveSess()
		e.ReloadWin("login")
	}, gwu.ETYPE_CLICK)
	header.Add(logout)
	win.Add(header)

	content := gwu.NewHorizontalPanel()
	content.SetVAlign(gwu.VA_TOP)
	content.Style().SetFullSize()

	showWrapper := gwu.NewPanel()
	selectedShow := gwu.NewLabel("")
	selectedShow.Style().SetFontWeight(gwu.FONT_WEIGHT_BOLD).SetFontSize("110%")
	showWrapper.Add(selectedShow)
	showWrapper.AddVSpace(10)

	links := gwu.NewPanel()
	links.Style().SetWhiteSpace(gwu.WHITE_SPACE_NOWRAP)

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

	links.Style().SetFullSize().SetBorderRight2(2, gwu.BRD_STYLE_SOLID, "#777777")
	l = gwu.NewLabel("Component Palette")
	l.Style().SetFontWeight(gwu.FONT_WEIGHT_BOLD).SetFontSize("120%").SetDisplay(gwu.DISPLAY_BLOCK).SetPaddingBottomPx(5)
	links.Add(l)
	l = gwu.NewLabel("Input components")
	l.Style().SetFontWeight(gwu.FONT_WEIGHT_BOLD).SetDisplay(gwu.DISPLAY_BLOCK).SetPaddingTopPx(5)
	links.Add(l)
	addShowLink("CheckBox", buildCheckBoxShow)
	addShowLink("ListBox", buildListBoxShow)
	addShowLink("TextBox", buildTextBoxShow)
	addShowLink("PasswBox", buildPasswBoxShow)
	addShowLink("RadioButton", buildRadioButtonShow)
	addShowLink("SwitchButton", buildSwitchButtonShow)
	l = gwu.NewLabel("Containers")
	l.Style().SetFontWeight(gwu.FONT_WEIGHT_BOLD).SetDisplay(gwu.DISPLAY_BLOCK).SetPaddingTopPx(5)
	links.Add(l)
	addShowLink("Link (as Container)", buildLinkShow)
	addShowLink("Panel", buildPanelShow)
	addShowLink("Table", buildTableShow)
	addShowLink("TabPanel", buildTabPanelShow)
	addShowLink("Window", buildWindowShow)
	l = gwu.NewLabel("Other components")
	l.Style().SetFontWeight(gwu.FONT_WEIGHT_BOLD).SetDisplay(gwu.DISPLAY_BLOCK).SetPaddingTopPx(5)
	links.Add(l)
	addShowLink("Button", buildButtonShow)
	addShowLink("Html", buildHtmlShow)
	addShowLink("Image", buildImageShow)
	addShowLink("Label", buildLabelShow)
	addShowLink("Link", buildLinkShow2)
	links.AddVConsumer()
	content.Add(links)
	content.Add(showWrapper)
	content.CellFmt(showWrapper).Style().SetFullWidth()

	win.Add(content)
	win.CellFmt(content).Style().SetFullSize()

	footer := gwu.NewHorizontalPanel()
	footer.Style().SetFullWidth().SetBorderTop2(2, gwu.BRD_STYLE_SOLID, "#777777").SetWhiteSpace(gwu.WHITE_SPACE_NOWRAP)
	footer.AddHConsumer()
	l = gwu.NewLabel("This app is written in and showcases Gowut version " + gwu.GOWUT_VERSION + ".")
	l.Style().SetFontStyle(gwu.FONT_STYLE_ITALIC)
	footer.Add(l)
	footer.AddHSpace(10)
	link := gwu.NewLink("Visit Gowut Home page", "https://sites.google.com/site/gowebuitoolkit/")
	footer.Add(link)
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
	tb := gwu.NewTextBox("admin")
	table.Add(tb, 0, 1)
	table.Add(gwu.NewLabel("Password:"), 1, 0)
	pb := gwu.NewPasswBox("a")
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

	win.SetFocusedCompId(b.Id())

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
