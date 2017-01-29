package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"golang.org/x/net/html"
)

func main() {
	fmt.Println("\nHello world:Start\n")

	files, _ := ioutil.ReadDir("html")
	for _, f := range files {
		fmt.Println(f.Name())
		parseFile("D:\\work\\go\\src\\github.com\\readBlog\\html\\" + f.Name())
	}
}

func parseFile(f string) {
	content, err := os.Open(f)
	if err != nil {
		fmt.Println(err)
	}
	z := html.NewTokenizer(content)

	isTitle := false

	for {
		tt := z.Next()

		fmt.Printf("%s\n", tt.String())

		switch {
		case tt == html.ErrorToken:
			// End of the document, we're done
			return
		case tt == html.StartTagToken:
			t := z.Token()

			// Check if the token is an <a> tag
			isTitle = t.Data == "h1"
			fmt.Println("found h1")
			if !isTitle {
				continue
			}

			// Extract the title value, if there is one
			//ok, title := getTitle(t)
			//if !ok {
			//	continue
			//}

		case tt == html.TextToken:
			fmt.Printf("TTesttoken: %b\n", isTitle)
			if isTitle {
				t := z.Token()
				fmt.Println("Title: " + t.String())
				isTitle = false

			}
		}
	}

}

// Helper function to pull the href attribute from a Token
func getTitle(t html.Token) (ok bool, title string) {
	// Iterate over all of the Token's attributes until we find an "href"

	title = t.Data
	ok = true
	return
}
