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
	isContent := false

	for {
		tt := z.Next()

		//fmt.Printf("%s\n", tt.String())

		switch {
		case tt == html.ErrorToken:
			// End of the document, we're done
			return
		case tt == html.StartTagToken:
			t := z.Token()

			isTitle = isTitleFunc(t)
			//fmt.Printf("%s\n", t.Data)

			isContent = isContentFunc(t)

		case tt == html.TextToken:
			//fmt.Printf("TTesttoken: %b\n", isTitle)
			if isTitle {
				t := z.Token()
				fmt.Println("Title: " + t.String())
				isTitle = false
			} else if isContent {
				//t := z.Token()
				fmt.Println("Content: " + tt.String())
				isContent = false
			}
		}
	}

}

func isTitleFunc(t html.Token) (ok bool) {
	if t.Data != "h1" {
		return false
	}
	for _, a := range t.Attr {
		//fmt.Printf("Val: %s\n", a.Val)
		if a.Val == "entry-title" {
			return true
		}
	}
	return false
}

func isContentFunc(t html.Token) (ok bool) {
	if t.Data != "div" {
		return false
	}
	for _, a := range t.Attr {
		//fmt.Printf("Val: %s\n", a.Val)
		if a.Val == "entry-content" {
			return true
		}
	}
	return false
}
