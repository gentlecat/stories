package index

import (
	"bytes"
	"fmt"
	"go.roman.zone/publisher/reader"
	"go.roman.zone/publisher/writer"
	"html/template"
	"log"
	"path"
)

func GenerateIndexPage(stories *[]reader.Story, tpl *template.Template, outputDir string) {
	log.Println("Generating index page...")
	defer log.Println("Finished generating the index page!")

	var templateOutput bytes.Buffer

	type ListItem struct {
		Path  string
		Story *reader.Story
	}
	var items []ListItem

	for i, s := range *stories {
		items = append(items, ListItem{Path: s.Name, Story: &(*stories)[i]})
	}
	type PageData struct {
		Title   string
		Stories []ListItem
	}

	if err := tpl.ExecuteTemplate(&templateOutput, "base", PageData{
		Title:   "",
		Stories: items,
	}); err != nil {
		log.Fatalf("Failed to render index page: %v", err)
	}

	writer.WriteFile(path.Join(outputDir, fmt.Sprintf("index.html")), templateOutput.Bytes())
}
