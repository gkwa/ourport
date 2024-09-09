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
			sort.Slice(urls, func(i, j int) bool {
				return extractNumber(urls[i]) < extractNumber(urls[j])
			})
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

	funcMap := template.FuncMap{
		"add": func(a, b, c int) int {
			return a + b + c
		},
	}

	tmpl, err := template.New("page").Funcs(funcMap).Parse(string(tmplContent))
	if err != nil {
		return fmt.Errorf("failed to parse template: %w", err)
	}

	totalGroups := len(groupList)
	for i := 0; i < totalGroups; i += groupsPerPage {
		end := i + groupsPerPage
		if end > totalGroups {
			end = totalGroups
		}

		fileName := fmt.Sprintf("ourport-images-%02d-%03d.md", groupsPerPage, i/groupsPerPage+1)
		file, err := os.Create(fileName)
		if err != nil {
			return err
		}
		defer file.Close()

		err = tmpl.Execute(file, struct {
			Groups        []GroupInfo
			GroupsInFile  int
			TotalGroups   int
			StartingIndex int
		}{
			Groups:        groupList[i:end],
			GroupsInFile:  end - i,
			TotalGroups:   totalGroups,
			StartingIndex: i,
		})
		if err != nil {
			return err
		}

		fmt.Printf("Generated %s\n", fileName)
	}

	return nil
}
