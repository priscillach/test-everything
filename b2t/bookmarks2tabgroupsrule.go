package b2t

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

type URLMatch struct {
	Method string `json:"method"`
	Target string `json:"target"`
	Value  string `json:"value"`
}

type Rule struct {
	GroupName  string     `json:"groupName"`
	ID         string     `json:"id"`
	RuleName   string     `json:"ruleName"`
	URLMatches []URLMatch `json:"urlMatches"`
}

type Meta struct {
	Name    string `json:"name"`
	Version int    `json:"version"`
}

type RulesOutput struct {
	Meta  Meta            `json:"meta"`
	Rules map[string]Rule `json:"-"`
}

func (r RulesOutput) MarshalJSON() ([]byte, error) {
	type Alias RulesOutput
	var out struct {
		Meta Meta `json:"meta"`
		*Alias
	}
	out.Meta = r.Meta

	m := make(map[string]interface{})
	m["meta"] = r.Meta
	for k, v := range r.Rules {
		m[k] = v
	}

	return json.Marshal(m)
}

func extractHostname(rawURL string) string {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return ""
	}
	hostname := parsedURL.Hostname()
	if strings.HasPrefix(hostname, "www.") {
		hostname = hostname[4:]
	}
	return hostname
}

func generateRuleID() string {
	return fmt.Sprintf("rule-%s", strings.ToLower(randomString(8)))
}

func randomString(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyz0123456789"
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func convertBookmarksToRules(filename string) (RulesOutput, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return RulesOutput{}, err
	}

	doc, err := html.Parse(strings.NewReader(string(content)))
	if err != nil {
		return RulesOutput{}, err
	}

	rulesById := make(map[string]Rule)
	var folderStack []string
	directBookmarks := make(map[string]bool)

	var traverse func(*html.Node)
	traverse = func(n *html.Node) {
		if n.Type == html.ElementNode {
			switch n.Data {
			case "h3":
				if n.FirstChild != nil {
					folderName := n.FirstChild.Data
					folderStack = append(folderStack, folderName)
				}
			case "a":
				currentFolder := "Bookmarks Bar"
				if len(folderStack) > 0 {
					currentFolder = folderStack[len(folderStack)-1]
				}

				for _, attr := range n.Attr {
					if attr.Key == "href" {
						hostname := extractHostname(attr.Val)
						if hostname != "" {
							var rule Rule
							ruleId := ""

							for id, r := range rulesById {
								if r.GroupName == currentFolder {
									rule = r
									ruleId = id
									break
								}
							}

							if ruleId == "" {
								ruleId = generateRuleID()
								rule = Rule{
									GroupName:  currentFolder,
									ID:         ruleId,
									RuleName:   currentFolder,
									URLMatches: []URLMatch{},
								}
							}

							exists := false
							for _, match := range rule.URLMatches {
								if match.Value == hostname {
									exists = true
									break
								}
							}
							if !exists {
								rule.URLMatches = append(rule.URLMatches, URLMatch{
									Method: "includes",
									Target: "hostname",
									Value:  hostname,
								})
								rulesById[rule.ID] = rule
								directBookmarks[currentFolder] = true
							}
						}
						break
					}
				}
			}
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			traverse(c)
		}

		if n.Type == html.ElementNode && n.Data == "dl" {
			if len(folderStack) > 0 {
				folderStack = folderStack[:len(folderStack)-1]
			}
		}
	}

	traverse(doc)

	return RulesOutput{
		Rules: rulesById,
		Meta: Meta{
			Name:    "tab-groups-rules",
			Version: 1,
		},
	}, nil
}
