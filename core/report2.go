package core

import (
	"path/filepath"
	"sort"

	"github.com/gkwa/ourport/tutorial"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func RunReport2() error {
	links, err := FetchImageLinks()
	if err != nil {
		return err
	}
	Report2(links)
	return nil
}

func Report2(links []tutorial.GetImageLinksRow) {
	p := message.NewPrinter(language.English)

	groups := make(map[string][]string)
	for _, link := range links {
		parentDir := filepath.Dir(link.Url)
		groups[parentDir] = append(groups[parentDir], link.Url)
	}

	groupNames := make([]string, 0, len(groups))
	for group := range groups {
		groupNames = append(groupNames, group)
	}
	sort.Strings(groupNames)

	for i, group := range groupNames {
		p.Printf("%d. %s (%d links)\n", i+1, group, len(groups[group]))
	}

	p.Printf("\nTotal number of groups: %d\n", len(groups))
	p.Printf("Total number of links: %d\n", len(links))
}
