package handler

import (
	"inventory/types"
)

func AuthorsWithCds(items []types.Item) []string {
	authorsWithCdsMap := map[string]bool{}
	cds := []types.Item{}
	bookAuthorMap := map[string]bool{}

	// Populate the authorMap and cds slice.
	for _, item := range items {

		if item.Type == "book" {
			// Add to the author map.
			bookAuthorMap[item.Author] = true
		} else if item.Type == "cd" {
			// Add to cds.
			cds = append(cds, item)
		}
	}

	// Loop through the cds checking if the author is in the book author map.
	for _, cd := range cds {
		if _, ok := bookAuthorMap[cd.Author]; ok {
			// This cd author is also a book author.
			authorsWithCdsMap[cd.Author] = true
		}
	}

	// Return the list of authorsWithCds key.
	authors := []string{}

	for author := range authorsWithCdsMap {
		authors = append(authors, author)
	}

	return authors
}
