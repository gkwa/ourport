package core

import (
	"path/filepath"

	"github.com/gkwa/ourport/tutorial"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func RunReport1() error {
	links, err := FetchImageLinks()
	if err != nil {
		return err
	}
	Report1(links)
	return nil
}

func Report1(links []tutorial.GetImageLinksRow) {
	p := message.NewPrinter(language.English)

	groups := make(map[string][]string)
	for _, link := range links {
		parentDir := filepath.Dir(link.Url)
		groups[parentDir] = append(groups[parentDir], link.Url)
	}

	p.Printf("Number of groups: %d\n", len(groups))
	p.Printf("Total number of links: %d\n", len(links))
}
