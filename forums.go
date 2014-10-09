pfackage main

import (
	"os"
	"fmt"
	"strings"
	"encoding/json"
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

func LoadForums() {
	fmt.Println("[INFO] forums.json loading")

	if len(os.Args) != 2 {
		fmt.Println("[ERROR] Bad usage")
		fmt.Println("Usage: ./shfr \"$(cat forums.json)\"")
		os.Exit(1)
	}

	data := os.Args[1]
	err := json.Unmarshal([]byte(data), &forum)
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

	fmt.Println("[INFO] forums.json loaded")
}
