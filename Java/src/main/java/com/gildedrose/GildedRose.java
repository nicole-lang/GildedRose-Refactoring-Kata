package com.gildedrose;

class GildedRose {
    Item[] items;

    public GildedRose(Item[] items) {
        this.items = items;
    }

    public void updateQuality() {
        for (int i = 0; i < items.length; i++) {
            Item item = items[i];
            String itemName = item.name;

            if (itemName.equals("Sulfuras, Hand of Ragnaros")) {
                continue;
            }

            if (isAgedProduct(itemName) && item.quality < 50) {
                addQuality(item);
            } else if (item.quality > 0) {
                // Conjured items degrade in quality twice as fast
                if (itemName.contains("Conjured")) {
                    item.quality--;
                }
                item.quality--;
            }

            item.sellIn--;

            if (item.sellIn < 0 && itemName.equals("Backstage passes to a TAFKAL80ETC concert")) {
                item.quality = 0;
            }
        }
    }

    public boolean isAgedProduct(String itemName) {
        return (itemName.equals("Aged Brie")
                || itemName.equals("Backstage passes to a TAFKAL80ETC concert"));
    }

    public void addQuality(Item item) {
        if (item.name.equals("Aged Brie")) {

            item.quality++;

        } else if (item.name.equals("Backstage passes to a TAFKAL80ETC concert")) {
            item.quality++;
            if (item.sellIn < 11 && item.quality < 50)
                item.quality++;

            if (item.sellIn < 6 && item.quality < 50)
                item.quality++;
        }
    }

}
