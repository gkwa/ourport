package core

import (
	"fmt"
	"net/url"
	"path"
	"sort"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func RunReport5() error {
	links, err := FetchImageLinks()
	if err != nil {
		return err
	}

	groups := make(map[string][]string)
	for _, link := range links {
		u, err := url.Parse(link.Url)
		if err != nil {
			continue
		}
		group := fmt.Sprintf("%s://%s%s", u.Scheme, u.Host, path.Dir(u.Path))
		groups[group] = append(groups[group], link.Url)
	}

	type groupInfo struct {
		name  string
		count int
	}

	var groupList []groupInfo
	for name, urls := range groups {
		if len(urls) > 3 {
			groupList = append(groupList, groupInfo{name, len(urls)})
		}
	}

	sort.Slice(groupList, func(i, j int) bool {
		return groupList[i].count < groupList[j].count
	})

	p := message.NewPrinter(language.English)
	totalGroups := len(groupList)
	p.Printf("Total number of groups: %d\n\n", totalGroups)

	for i, group := range groupList {
		p.Printf("Group %d: %s (Count: %d)\n", i+1, group.name, group.count)

		sort.Slice(groups[group.name], func(i, j int) bool {
			numI := extractNumber(groups[group.name][i])
			numJ := extractNumber(groups[group.name][j])
			return numI < numJ
		})

		for _, url := range groups[group.name] {
			fmt.Printf("  %s\n", url)
		}
		fmt.Println()
	}

	return nil
}
