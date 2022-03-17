package handler

import (
	"inventory/types"
	"sort"
)

func MostExpensive(items []types.Item, itemType string) []types.Item {
	list := []types.Item{}

	sort.Slice(items, func(i, j int) bool { return items[i].Price > items[j].Price })

	for _, item := range items {

		if item.Type == itemType {
			list = append(list, item)
		}

		if len(list) >= 5 {
			// We have 5 items in the list.
			break
		}
	}

	return list
}
