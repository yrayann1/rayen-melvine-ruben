package main

import (
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Artist struct {
	Name         string
	Members      []string
	CreationDate int
	FirstAlbum   int
	Locations    []string
}

var artists = []Artist{
	{
		Name:         "Queen",
		Members:      []string{"Freddie Mercury", "Brian May"},
		CreationDate: 1970,
		FirstAlbum:   1973,
		Locations:    []string{"London", "Paris"},
	},
	{
		Name:         "Metallica",
		Members:      []string{"James Hetfield", "Lars Ulrich"},
		CreationDate: 1981,
		FirstAlbum:   1983,
		Locations:    []string{"Los Angeles", "Berlin"},
	},
}

func containsIgnoreCase(a, b string) bool {
	return strings.Contains(strings.ToLower(a), strings.ToLower(b))
}

func getSuggestions(input string) []string {
	if input == "" {
		return []string{}
	}

	var res []string
	for _, a := range artists {
		if containsIgnoreCase(a.Name, input) {
			res = append(res, a.Name+" - artist/band")
		}
		for _, m := range a.Members {
			if containsIgnoreCase(m, input) {
				res = append(res, m+" - member")
			}
		}
		for _, l := range a.Locations {
			if containsIgnoreCase(l, input) {
				res = append(res, l+" - location")
			}
		}
		if containsIgnoreCase(strconv.Itoa(a.CreationDate), input) {
			res = append(res, strconv.Itoa(a.CreationDate)+" - creation date")
		}
		if containsIgnoreCase(strconv.Itoa(a.FirstAlbum), input) {
			res = append(res, strconv.Itoa(a.FirstAlbum)+" - first album")
		}
	}
	return res
}

func main() {
	myApp := app.New()
	w := myApp.NewWindow("Groupie Tracker - Search")

	search := widget.NewEntry()
	search.SetPlaceHolder("Search artist, member, location...")

	suggestions := widget.NewList(
		func() int { return 0 },
		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {},
	)

	search.OnChanged = func(text string) {
		data := getSuggestions(text)
		suggestions.Length = func() int {
			return len(data)
		}
		suggestions.UpdateItem = func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(data[i])
		}
		suggestions.Refresh()
	}

	w.SetContent(container.NewVBox(
		search,
		suggestions,
	))

	w.Resize(fyne.NewSize(500, 400))
	w.ShowAndRun()
}
