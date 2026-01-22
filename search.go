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



func applyFilters(minC, maxC, minA, maxA int, membersCount []int, location string) []Artist {
	var result []Artist

	for _, a := range artists {
		if a.CreationDate < minC || a.CreationDate > maxC {
			continue
		}
		if a.FirstAlbum < minA || a.FirstAlbum > maxA {
			continue
		}

		if len(membersCount) > 0 {
			valid := false
			for _, n := range membersCount {
				if len(a.Members) == n {
					valid = true
				}
			}
			if !valid {
				continue
			}
		}

		if location != "" {
			found := false
			for _, l := range a.Locations {
				if l == location {
					found = true
				}
			}
			if !found {
				continue
			}
		}

		result = append(result, a)
	}

	return result
}

func main() {
	myApp := app.New()
	w := myApp.NewWindow("Groupie Tracker")


	search := widget.NewEntry()
	search.SetPlaceHolder("Search artist, member, location...")

	suggestions := widget.NewList(
		func() int { return 0 },
		func() fyne.CanvasObject { return widget.NewLabel("") },
		func(i widget.ListItemID, o fyne.CanvasObject) {},
	)

	search.OnChanged = func(text string) {
		data := getSuggestions(text)
		suggestions.Length = func() int { return len(data) }
		suggestions.UpdateItem = func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(data[i])
		}
		suggestions.Refresh()
	}



	minCreation := widget.NewSlider(1950, 2025)
	maxCreation := widget.NewSlider(1950, 2025)
	minAlbum := widget.NewSlider(1950, 2025)
	maxAlbum := widget.NewSlider(1950, 2025)

	minCreation.SetValue(1950)
	maxCreation.SetValue(2025)
	minAlbum.SetValue(1950)
	maxAlbum.SetValue(2025)

	c2 := widget.NewCheck("2 members", nil)
	c3 := widget.NewCheck("3 members", nil)
	c4 := widget.NewCheck("4 members", nil)

	location := widget.NewSelect(
		[]string{"London", "Paris", "Berlin", "Los Angeles"},
		nil,
	)

	apply := widget.NewButton("Apply Filters", func() {
		var members []int
		if c2.Checked {
			members = append(members, 2)
		}
		if c3.Checked {
			members = append(members, 3)
		}
		if c4.Checked {
			members = append(members, 4)
		}

		_ = applyFilters(
			int(minCreation.Value),
			int(maxCreation.Value),
			int(minAlbum.Value),
			int(maxAlbum.Value),
			members,
			location.Selected,
		)
	})

	filters := container.NewVBox(
		widget.NewLabel("Creation Date"),
		minCreation, maxCreation,
		widget.NewLabel("First Album Date"),
		minAlbum, maxAlbum,
		widget.NewLabel("Members"),
		c2, c3, c4,
		widget.NewLabel("Location"),
		location,
		apply,
	)

	w.SetContent(container.NewVBox(
		search,
		suggestions,
		filters,
	))

	w.Resize(fyne.NewSize(500, 650))
	w.ShowAndRun()
}
