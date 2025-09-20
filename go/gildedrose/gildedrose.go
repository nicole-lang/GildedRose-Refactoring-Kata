package gildedrose

import (
	"strings"
)

type Item struct {
	Name            string
	SellIn, Quality int
}

func UpdateQuality(items []*Item) {
	for _, item := range items {
		var shiftAmt int

		//Update shift amount based on Name
		switch item.Name {
		case "Sulfuras, Hand of Ragnaros":
			continue
		case "Aged Brie":
			shiftAmt = 1

		case "Backstage passes to a TAFKAL80ETC concert":
			{
				// Backstage passes gain quality as the concert date approaches and fall to 0 after the concert
				if item.SellIn > 0 && item.SellIn < 11 {
					if item.SellIn < 6 {
						shiftAmt = 3
					} else {
						shiftAmt = 2
					}
				} else {
					item.Quality = 0
				}
			}

		default:
			{
				if item.Quality > 0 {
					if strings.Contains(item.Name, "Conjured") {
						//Conjured items degrade in quality twice as fast
						shiftAmt = -2
					} else {
						shiftAmt = -1
					}
				}
			}
		}

		if item.SellIn > 0 {
			item.SellIn--
		}
		if item.SellIn == 0 && item.Name == "Backstage passes to a TAFKAL80ETC concert" {
			item.Quality = 0
			continue
		}

		if item.SellIn < 0 {
			if item.Name != "Aged Brie" && item.Quality > 0 {
				shiftAmt--
			}
		}

		item.Quality += shiftAmt

		if item.Quality > 50 {
			item.Quality = 50
		}
	}
}
