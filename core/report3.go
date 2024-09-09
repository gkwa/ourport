package core

import (
	"fmt"

	"github.com/gkwa/ourport/tutorial"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func RunReport3() error {
	links, err := FetchImageLinks()
	if err != nil {
		return err
	}
	Report3(links)
	return nil
}

func Report3(links []tutorial.GetImageLinksRow) {
	p := message.NewPrinter(language.English)

	fmt.Println("All image links:")
	for i, link := range links {
		p.Printf("%d. %s\n", i+1, link.Url)
	}

	totalLinks := len(links)
	p.Printf("\nTotal number of image links: %d\n", totalLinks)
}
