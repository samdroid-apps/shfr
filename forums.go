package main

import (
	"fmt"
	"strings"
	"io/ioutil"
	"encoding/json"
	"github.com/andreaskoch/go-fswatch"
)

const (
	FORUMS_FILE string = "forums.json"
)

type ForumsFile struct {
	Forums map[string]string `json:"forums"`
	URL    string            `json:"url"`
}

var (
	forum ForumsFile
)

func GetInfo() ForumsFile {
	return forum
}

func GetForum(id string) (string, bool) {
	name, found := forum.Forums[id]
	if !found {
		return forum.URL + "/t/category-not-found/", true;
	} else {
		return forum.URL + "/category/" + name, false;
	}
}

func loadForums() {
	fmt.Println("[INFO] forums.json changed, reloading")

	data, err := ioutil.ReadFile(FORUMS_FILE)
	if err != nil {
		fmt.Println("[ERROR] Couldn't load forums.json - " +
					"file error (did you forget to add it?)")
		return
	}

	err = json.Unmarshal(data, &forum)
	if err != nil {
		fmt.Println("[ERROR] Couldn't load forums.json - " +
					"json load error")
		return
	}

	if !(strings.HasPrefix(forum.URL, "http://") ||
		strings.HasPrefix(forum.URL, "https://")) {
		forum.URL = "http://" + forum.URL
	}

	if strings.HasSuffix(forum.URL, "/") {
		forum.URL = forum.URL[:len(forum.URL)-1]
	}

	fmt.Println("[INFO] forums.json reloaded")
}

func WatchForums() {
	loadForums()
	fw := fswatch.NewFileWatcher(FORUMS_FILE).Start()
	for fw.IsRunning() {
		select {
		case <-fw.Modified:
			loadForums()
		default:
		}
	}
}
