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

func parseHTMLCode(fileName string) []byte {

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

func extractLinks(htmlData *html.Node) []*html.Node {
	if htmlData.Type == html.ElementNode && htmlData.Data == "a" {
		return []*html.Node{htmlData}
	}
	var htmlLinks []*html.Node
	for c := htmlData.FirstChild; c != nil; c = c.NextSibling {
		htmlLinks = append(htmlLinks, extractLinks(c)...)
	}
	return htmlLinks
}

func buildLinks(n *html.Node) Link {
	var l Link
	for _, a := range n.Attr {
		if a.Key == "href" {
			l.Href = a.Val
			break
		}
	}
	l.Text = getTextFromNode(n)
	return l
}

func getTextFromNode(builtLink *html.Node) string {
	if builtLink.Type == html.TextNode {
		return builtLink.Data
	}
	if builtLink.Type != html.ElementNode {
		return ""
	}
	var ret string
	for c := builtLink.FirstChild; c != nil; c = c.NextSibling {
		ret += getTextFromNode(c)
	}
	return strings.Join(strings.Fields(ret), " ")
}

// ParserPipeline runs the whole pipeline in one go
func ParserPipeline(filename string) []Link {
	fmt.Println("*Reading Data from template file...")
	templateString := parseHTMLCode(filename)
	fmt.Println("**Getting Tree structure from HTML...")
	templateHTMLNode := getHTMLTree(templateString)
	fmt.Println("***Extracting all the links from html tree...")
	fmt.Println("**************************************")
	n := extractLinks(templateHTMLNode)
	var links []Link
	for _, i := range n {
		links = append(links, buildLinks(i))
	}
	return links
}
