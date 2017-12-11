package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
)

var watcher *fsnotify.Watcher

func main() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Println("Error, ", err)
	}
	defer watcher.Close()

	if err := filepath.Walk("/home/joker/go/src/dir_sync/test_dir/", watchDir); err != nil {
		fmt.Println("Error, ", err)
	}

	done := make(chan bool)

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

	<-done

}

func watchDir(path string, fi os.FileInfo, err error) error {
	if fi.Mode().IsDir() {
		err = watcher.Add(path)
		fmt.Println("watchDir Error, ", err)
	}
	return err
}
