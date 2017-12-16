package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
)

func main() {
	mainPath := "/home/joker/go/src/github.com/navinds25/DirSync/test_dir"

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Println("Error, ", err)
	}
	defer watcher.Close()

	//open the directory file
	dir, err := os.Open(mainPath)
	if err != nil {
		fmt.Printf("Error while opening mainPath %s, Error %s \n", mainPath, err)
	}
	defer dir.Close()

	// to read the contents of the directory file
	mainContent, err := dir.Readdir(-1)
	if err != nil {
		fmt.Printf("Error listing contents of %s , Exception: %s \n", mainPath, err)
	}

	// Add the main path to the watchlist
	if err = watcher.Add(mainPath); err != nil {
		fmt.Printf("Error adding main dir %s to watchlist Error: %s \n", mainPath, err)
	}

	// looping through contents of main dir to find sub directories
	for _, content := range mainContent {
		if content.Mode().IsDir() {
			path := filepath.Join(mainPath, content.Name())
			fmt.Println("Found Dir", path)
			// add it to watcher
			if err := watcher.Add(path); err != nil {
				fmt.Printf("Error adding dir %s Error: %s \n", path, err)
			}
		}
	}

	// channel+ go routine + anonymous func for watching directories
	dirwatch := make(chan bool)

	go func() {
		for {
			select {
			case event := <-watcher.Events:
				fmt.Println("Event! ", event)
			case err := <-watcher.Errors:
				fmt.Println("Error!", err)
			}
		}
	}()

	<-dirwatch
}
