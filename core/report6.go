package core

import (
	"embed"
	"fmt"
	"net/url"
	"os"
	"path"
	"sort"
	"text/template"
)

//go:embed templates
var templatesFS embed.FS

func RunReport6(groupsPerPage int) error {
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

	type GroupInfo struct {
		Name  string
		Count int
		URLs  []string
	}

	var groupList []GroupInfo
	for name, urls := range groups {
		if len(urls) > 1 {
			groupList = append(groupList, GroupInfo{Name: name, Count: len(urls), URLs: urls})
		}
	}

	sort.Slice(groupList, func(i, j int) bool {
		return groupList[i].Count > groupList[j].Count
	})

	tmplContent, err := templatesFS.ReadFile("templates/report6.md.tmpl")
	if err != nil {
		return fmt.Errorf("failed to read template: %w", err)
	}

	tmpl, err := template.New("page").Parse(string(tmplContent))
	if err != nil {
		return fmt.Errorf("failed to parse template: %w", err)
	}

	for i := 0; i < len(groupList); i += groupsPerPage {
		end := i + groupsPerPage
		if end > len(groupList) {
			end = len(groupList)
		}

		fileName := fmt.Sprintf("ourport-images-%03d.md", i/groupsPerPage+1)
		file, err := os.Create(fileName)
		if err != nil {
			return err
		}
		defer file.Close()

		err = tmpl.Execute(file, struct {
			Groups []GroupInfo
		}{
			Groups: groupList[i:end],
		})
		if err != nil {
			return err
		}

		fmt.Printf("Generated %s\n", fileName)
	}

	return nil
}
