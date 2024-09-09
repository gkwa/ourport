package core

import (
	"fmt"
	"math/rand"
	"path/filepath"
	"sort"
	"time"

	"github.com/gkwa/ourport/tutorial"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func RunReport4() error {
	links, err := FetchImageLinks()
	if err != nil {
		return err
	}
	Report4(links)
	return nil
}

func Report4(links []tutorial.GetImageLinksRow) {
	p := message.NewPrinter(language.English)

	groups := make(map[string][]string)
	for _, link := range links {
		parentDir := filepath.Dir(link.Url)
		groups[parentDir] = append(groups[parentDir], link.Url)
	}

	// Convert map to slice for randomization
	var groupSlice []struct {
		name  string
		links []string
	}
	for name, links := range groups {
		groupSlice = append(groupSlice, struct {
			name  string
			links []string
		}{name, links})
	}

	// Create a new random source and shuffle the slice
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	rng.Shuffle(len(groupSlice), func(i, j int) {
		groupSlice[i], groupSlice[j] = groupSlice[j], groupSlice[i]
	})

	// Print randomized groups and their links
	for i, group := range groupSlice {
		p.Printf("%d. Group: %s (%d links)\n", i+1, group.name, len(group.links))

		// Sort links within the group
		sort.Slice(group.links, func(i, j int) bool {
			return extractNumber(group.links[i]) < extractNumber(group.links[j])
		})

		for j, link := range group.links {
			p.Printf("   %d.%d. %s\n", i+1, j+1, link)
		}
		fmt.Println() // Add a blank line between groups
	}

	p.Printf("Total number of groups: %d\n", len(groups))
	p.Printf("Total number of links: %d\n", len(links))
}
