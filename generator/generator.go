package main

import (
	"bytes"
	"fmt"
	"gopkg.in/yaml.v2"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

const (
	govYamlFile         = "gov.yaml"
	indexFilename       = "README.md"
	beginCustomMarkdown = "<!-- BEGIN CUSTOM CONTENT -->"
	endCustomMarkdown   = "<!-- END CUSTOM CONTENT -->"
)

// GitHubTeam represents a GitHub team reference.
type GitHubTeam struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description,omitempty"`
}

// Term holds start and end dates.
type Term struct {
	Start string `yaml:"start,omitempty"`
	End   string `yaml:"end,omitempty"`
}

// Person holds person data.
type Person struct {
	Name    string `yaml:"name"`
	GitHub  string `yaml:"github"`
	Slack   string `yaml:"slack,omitempty"`
	Seat    string `yaml:"seat,omitempty"`
	Role    string `yaml:"role,omitempty"`
	Company string `yaml:"company,omitempty"`
	Term    Term   `yaml:"term,omitempty"`
}

// Meeting holds meeting data.
type Meeting struct {
	Description     string `yaml:"description"`
	RecordingsURL   string `yaml:"recordings_url,omitempty"`
	MeetingURL      string `yaml:"meeting_url,omitempty"`
	MeetingNotesURL string `yaml:"meeting_notes_url,omitempty"`
}

// Contact holds contact information.
type Contact struct {
	Slack        string       `yaml:"slack,omitempty"`
	SlackChannel string       `yaml:"slack_channel,omitempty"`
	MailingList  string       `yaml:"mailing_list,omitempty"`
	GitHubTeams  []GitHubTeam `yaml:"github_teams,omitempty"`
	Liaison      []Person     `yaml:"liaison,omitempty"`
}

// Group represents a TAB group or User Group.
type Group struct {
	Dir              string    `yaml:"dir"`
	Name             string    `yaml:"name"`
	Leadership       []Person  `yaml:"leadership,omitempty"`
	Members          []Person  `yaml:"members,omitempty"`
	Emeritus         []Person  `yaml:"emeritus,omitempty"`
	MissionStatement string    `yaml:"mission_statement,omitempty"`
	Meetings         []Meeting `yaml:"meetings,omitempty"`
	Contact          Contact   `yaml:"contact,omitempty"`
	CharterLink      string    `yaml:"charter_link,omitempty"`
	Label            string    `yaml:"label,omitempty"`
}

// Config is the top-level governance configuration.
type Config struct {
	TAB        []Group `yaml:"tab"`
	UserGroups []Group `yaml:"user_groups"`
}

// TemplateData extends Group with computed fields for template rendering.
type TemplateData struct {
	Group
	GroupType   string // "tab" or "user-group"
	RepoBaseURL string
}

var (
	templateDir = "."
)

func main() {
	configPath := filepath.Join("..", govYamlFile)

	data, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatalf("Failed to read %s: %v", configPath, err)
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		log.Fatalf("Failed to parse YAML: %v", err)
	}

	// Load templates.
	funcMap := template.FuncMap{
		"lower":     strings.ToLower,
		"replace":   strings.ReplaceAll,
		"trimSpace": strings.TrimSpace,
		"githubLink": func(p Person) string {
			if p.GitHub != "" {
				return fmt.Sprintf("[%s](https://github.com/%s)", p.Name, p.GitHub)
			}
			return p.Name
		},
	}

	groupTmpl := loadTemplate("group_readme.tmpl", funcMap)

	repoBaseURL := "https://github.com/cncf/tab"

	// Process TAB groups.
	processGroups(config.TAB, "tab", "..", repoBaseURL, groupTmpl)

	// Process User Groups.
	processGroups(config.UserGroups, "user-group", filepath.Join("..", "user-groups"), repoBaseURL, groupTmpl)

	log.Println("README files have been generated successfully.")
}

func processGroups(groups []Group, groupType, baseDir, repoBaseURL string, groupTmpl *template.Template) {
	if err := ensureDir(baseDir); err != nil {
		log.Fatalf("Failed to create base directory %s: %v", baseDir, err)
	}

	for _, group := range groups {
		if group.Dir == "" {
			group.Dir = slugify(group.Name)
		}

		groupDir := filepath.Join(baseDir, group.Dir)
		if err := ensureDir(groupDir); err != nil {
			log.Fatalf("Failed to create directory for %s: %v", group.Name, err)
		}

		td := TemplateData{
			Group:       group,
			GroupType:   groupType,
			RepoBaseURL: repoBaseURL,
		}

		if err := writeTemplate(groupTmpl, td, filepath.Join(groupDir, indexFilename)); err != nil {
			log.Fatalf("Failed to generate README for %s: %v", group.Name, err)
		}
	}
}

func loadTemplate(name string, funcMap template.FuncMap) *template.Template {
	path := filepath.Join(templateDir, name)
	content, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Failed to read template %s: %v", path, err)
	}
	tmpl, err := template.New(name).Funcs(funcMap).Parse(string(content))
	if err != nil {
		log.Fatalf("Failed to parse template %s: %v", name, err)
	}
	return tmpl
}

func writeTemplate(tmpl *template.Template, data interface{}, outPath string) error {
	customContent, err := getExistingCustomContent(outPath)
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return fmt.Errorf("template execution failed for %s: %w", outPath, err)
	}

	buf.WriteString(beginCustomMarkdown)
	if customContent == "" {
		buf.WriteByte('\n')
	} else {
		buf.WriteString(customContent)
	}
	buf.WriteString(endCustomMarkdown)
	buf.WriteByte('\n')

	return os.WriteFile(outPath, buf.Bytes(), 0o644)
}

func getExistingCustomContent(path string) (string, error) {
	content, err := os.ReadFile(path)
	if os.IsNotExist(err) {
		return "", nil
	}
	if err != nil {
		return "", fmt.Errorf("read existing custom content from %s: %w", path, err)
	}

	beginCount := strings.Count(string(content), beginCustomMarkdown)
	endCount := strings.Count(string(content), endCustomMarkdown)
	if beginCount == 0 && endCount == 0 {
		return "", nil
	}
	if beginCount != 1 || endCount != 1 {
		return "", fmt.Errorf("invalid custom content markers in %s: expected one begin and one end marker", path)
	}

	begin := strings.Index(string(content), beginCustomMarkdown)
	end := strings.Index(string(content), endCustomMarkdown)
	if end < begin {
		return "", fmt.Errorf("invalid custom content markers in %s: end marker appears before begin marker", path)
	}

	return string(content[begin+len(beginCustomMarkdown) : end]), nil
}

func slugify(name string) string {
	s := strings.ToLower(name)
	s = strings.ReplaceAll(s, " ", "-")
	s = strings.ReplaceAll(s, "/", "-")
	return s
}

func ensureDir(dirPath string) error {
	return os.MkdirAll(dirPath, 0o755)
}
