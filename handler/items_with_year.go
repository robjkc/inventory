package handler

import (
	"inventory/types"
	"regexp"
)

func ItemsWithYear(items []types.Item) []types.Item {
	list := []types.Item{}
	// Map of titles we have already processed.
	titleMap := map[string]bool{}

	yearRegex := regexp.MustCompile(`\d\d\d\d`)

	for _, item := range items {
		_, ok := titleMap[item.Title]
		if ok {
			// We already found this title.
			continue
		}

		// Check if title is a year.
		if !hasTitle(item, titleMap) && yearRegex.MatchString(item.Title) {
			list = append(list, item)
			titleMap[item.Title] = true
		}

		if item.Type == "cd" {
			// Check if track name is a year.
			for _, track := range item.Tracks {
				if !hasTitle(item, titleMap) && yearRegex.MatchString(track.Name) {
					list = append(list, item)
					titleMap[item.Title] = true
				}
			}
		}
		if item.Type == "book" {
			// Check if book chapter name is a year.
			for _, chapter := range item.Chapters {
				if !hasTitle(item, titleMap) && yearRegex.MatchString(chapter) {
					list = append(list, item)
					titleMap[item.Title] = true
				}
			}
		}
	}

	return list
}

func hasTitle(item types.Item, titleMap map[string]bool) bool {
	_, ok := titleMap[item.Title]
	return ok
}
