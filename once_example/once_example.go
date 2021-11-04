package main

import (
	"log"
	"sync"
)

func main() {
	log.Println(Icon("clubs"))
}

var icons map[string]string

func loadIcons() {
	icons = map[string]string{
		"spades": "spades.png",
		"hearts": "hearts.png",
		"clubs":  "clubs.png",
	}
}

var loadIconsOnce sync.Once

func Icon(name string) string {
	loadIconsOnce.Do(loadIcons)
	return icons[name]
}
