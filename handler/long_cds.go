package handler

import (
	"inventory/types"
)

func LongCds(items []types.Item) []types.Item {
	list := []types.Item{}

	for _, item := range items {

		if item.Type == "cd" {
			totalSeconds := 0
			tracks := item.Tracks
			for _, track := range tracks {
				totalSeconds += track.Seconds
			}

			if totalSeconds > 3600 {
				list = append(list, item)
			}
		}
	}

	return list
}
