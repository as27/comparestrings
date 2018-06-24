package main

import (
	"strings"

	"github.com/as27/govuegui"
	"github.com/as27/govuegui/gui/photon"
)

func main() {
	gui := govuegui.NewGui(photon.Template)
	gui.Title = "Compare Strings"

	inBox := gui.Form("main").Box("Input")

	inBox.Textarea("String1").Set("")
	inBox.Textarea("String2").Set("")

	outBox := gui.Form("main").Box("Results")
	outBox.List("Not in 1").Set([]string{})
	outBox.List("Not in 2").Set([]string{})

	inBox.Button("Compare strings").Action(func(gui *govuegui.Gui) {
		s1 := inBox.Textarea("String1").Get().(string)
		s2 := inBox.Textarea("String2").Get().(string)
		outBox.List("Not in 1").Set(compareStrings(s1, s2))
		outBox.List("Not in 2").Set(compareStrings(s2, s1))
		gui.Update()
		outBox.Active()
	})
	govuegui.Serve(gui)
}

func compareStrings(s1, s2 string) []string {
	var out []string
	sl1 := parseString(s1)
	sl2 := parseString(s2)
	for _, s := range sl2 {
		if !stringInSlice(s, sl1) {
			out = append(out, s)
		}
	}
	return out
}

func parseString(s string) []string {
	var out []string
	lines := strings.Split(s, "\n")
	for _, l := range lines {
		ts := strings.TrimSpace(l)
		if ts == "" {
			continue
		}
		out = append(out, ts)
	}
	return out
}

func stringInSlice(search string, slice []string) bool {
	for _, s := range slice {
		if search == s {
			return true
		}
	}
	return false
}
