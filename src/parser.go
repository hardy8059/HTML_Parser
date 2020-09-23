package src

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	html "golang.org/x/net/html"
)

// Link Represents a json object of a hyperlink with link and text fields...
type Link struct {
	Href string `json:"href"`
	Text string `json:"text"`
}

func getTemplate(pathToTemplate string, filename string) ([]byte, error) {
	return ioutil.ReadFile(filepath.Join(pathToTemplate, filename))
}

func parseHTMLCode() []byte {
	fileName := "index.html"

	paths := initialisePaths()

	temp, err := getTemplate(paths.pathToTemplates, fileName)
	if err != nil {
		panic(err)
	}
	return temp
}

func getHTMLTree(fileBytes []byte) *html.Node {
	templateHTMLTree, err := html.Parse(strings.NewReader(string(fileBytes)))
	if err != nil {
		panic(err)
	}
	return templateHTMLTree
}

func extractLinks(htmlData *html.Node) {
	if htmlData.Type == html.ElementNode {
		fmt.Println(htmlData.Data)
		fmt.Println(htmlData.Parent)
	}
	for c := htmlData.FirstChild; c != nil; c = c.NextSibling {
		extractLinks(c)
	}
}

// ParserPipeline runs the whole pipeline in one go
func ParserPipeline() {
	fmt.Println("Reading Data from template file...")
	templateString := parseHTMLCode()
	fmt.Println("Getting Tree structure from HTML...")
	templateHTMLNode := getHTMLTree(templateString)
	fmt.Println("Extracting all the links from go...")
	extractLinks(templateHTMLNode)
}
