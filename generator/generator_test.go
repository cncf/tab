package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"text/template"

	"gopkg.in/yaml.v2"
)

// ---------------------------------------------------------------------------
// slugify
// ---------------------------------------------------------------------------

func TestSlugify(t *testing.T) {
	tests := []struct {
		input, want string
	}{
		{"Developer Experience", "developer-experience"},
		{"Public Sector", "public-sector"},
		{"Already-Lowercase", "already-lowercase"},
		{"With / Slash", "with---slash"},
		{"simple", "simple"},
		{"", ""},
	}
	for _, tt := range tests {
		if got := slugify(tt.input); got != tt.want {
			t.Errorf("slugify(%q) = %q, want %q", tt.input, got, tt.want)
		}
	}
}

// ---------------------------------------------------------------------------
// ensureDir
// ---------------------------------------------------------------------------

func TestEnsureDir(t *testing.T) {
	dir := t.TempDir()
	nested := filepath.Join(dir, "a", "b", "c")
	if err := ensureDir(nested); err != nil {
		t.Fatalf("ensureDir failed: %v", err)
	}
	info, err := os.Stat(nested)
	if err != nil {
		t.Fatalf("directory not created: %v", err)
	}
	if !info.IsDir() {
		t.Fatal("expected directory")
	}

	// idempotent
	if err := ensureDir(nested); err != nil {
		t.Fatalf("ensureDir not idempotent: %v", err)
	}
}

// ---------------------------------------------------------------------------
// YAML parsing (Config)
// ---------------------------------------------------------------------------

func TestConfigParsing(t *testing.T) {
	yamlData := `
tab:
  - name: Test TAB
    dir: test-tab
    mission_statement: Test mission
    leadership:
      - name: Alice
        github: alice
        slack: alice-slack
        company: ACME
        role: Chair
    members:
      - name: MemberOne
        github: member1
        company: BigCo
      - name: MemberTwo
        github: member2
        slack: member2-slack
        company: SmallCo
    emeritus:
      - name: Bob
        github: bob
    meetings:
      - description: Weekly sync
        meeting_url: https://example.com/meet
        meeting_notes_url: https://example.com/notes
        recordings_url: https://example.com/recordings
    contact:
      slack: https://slack.example.com
      slack_channel: test-tab
      mailing_list: tab@example.com
      github_teams:
        - name: tab-team
          description: The TAB team
      liaison:
        - name: Charlie
          github: charlie
user_groups:
  - name: UG One
    dir: ug-one
    leadership: []
`
	var config Config
	if err := yaml.Unmarshal([]byte(yamlData), &config); err != nil {
		t.Fatalf("failed to parse YAML: %v", err)
	}

	if len(config.TAB) != 1 {
		t.Fatalf("expected 1 TAB group, got %d", len(config.TAB))
	}
	tab := config.TAB[0]
	if tab.Name != "Test TAB" {
		t.Errorf("tab name = %q, want %q", tab.Name, "Test TAB")
	}
	if tab.Dir != "test-tab" {
		t.Errorf("tab dir = %q, want %q", tab.Dir, "test-tab")
	}
	if len(tab.Leadership) != 1 {
		t.Fatalf("expected 1 leader, got %d", len(tab.Leadership))
	}
	if tab.Leadership[0].GitHub != "alice" {
		t.Errorf("leader github = %q", tab.Leadership[0].GitHub)
	}
	if tab.Leadership[0].Role != "Chair" {
		t.Errorf("leader role = %q, want Chair", tab.Leadership[0].Role)
	}
	if tab.Leadership[0].Slack != "alice-slack" {
		t.Errorf("leader slack = %q, want alice-slack", tab.Leadership[0].Slack)
	}
	if len(tab.Members) != 2 {
		t.Fatalf("expected 2 members, got %d", len(tab.Members))
	}
	if tab.Members[0].GitHub != "member1" {
		t.Errorf("member[0] github = %q", tab.Members[0].GitHub)
	}
	if tab.Members[1].Slack != "member2-slack" {
		t.Errorf("member[1] slack = %q, want member2-slack", tab.Members[1].Slack)
	}
	if len(tab.Emeritus) != 1 {
		t.Errorf("expected 1 emeritus, got %d", len(tab.Emeritus))
	}
	if tab.Emeritus[0].GitHub != "bob" {
		t.Errorf("emeritus github = %q", tab.Emeritus[0].GitHub)
	}
	if len(tab.Meetings) != 1 {
		t.Fatalf("expected 1 meeting, got %d", len(tab.Meetings))
	}
	if tab.Meetings[0].RecordingsURL != "https://example.com/recordings" {
		t.Errorf("recordings url wrong")
	}
	if len(tab.Contact.GitHubTeams) != 1 {
		t.Errorf("expected 1 github team")
	}
	if len(tab.Contact.Liaison) != 1 {
		t.Errorf("expected 1 liaison")
	}

	if len(config.UserGroups) != 1 {
		t.Fatalf("expected 1 user group, got %d", len(config.UserGroups))
	}
}

func TestConfigParsing_PersonSlackAndTerm(t *testing.T) {
	yamlData := `
tab:
  - name: Term Test
    dir: term-test
    leadership:
      - name: Leader
        github: leader
        slack: leader-slack
        company: Co
        term:
          start: "2025-01"
          end: "2027-01"
`
	var config Config
	if err := yaml.Unmarshal([]byte(yamlData), &config); err != nil {
		t.Fatalf("failed to parse YAML: %v", err)
	}
	p := config.TAB[0].Leadership[0]
	if p.Slack != "leader-slack" {
		t.Errorf("slack = %q, want leader-slack", p.Slack)
	}
	if p.Term.Start != "2025-01" {
		t.Errorf("term.start = %q, want 2025-01", p.Term.Start)
	}
	if p.Term.End != "2027-01" {
		t.Errorf("term.end = %q, want 2027-01", p.Term.End)
	}
}

func TestConfigParsing_EmptyFields(t *testing.T) {
	yamlData := `
tab:
  - name: Minimal
    dir: minimal
user_groups: []
`
	var config Config
	if err := yaml.Unmarshal([]byte(yamlData), &config); err != nil {
		t.Fatalf("failed to parse YAML: %v", err)
	}
	if len(config.TAB) != 1 {
		t.Fatalf("expected 1 TAB group, got %d", len(config.TAB))
	}
	tab := config.TAB[0]
	if tab.Leadership != nil {
		t.Errorf("expected nil leadership, got %v", tab.Leadership)
	}
	if tab.Members != nil {
		t.Errorf("expected nil members, got %v", tab.Members)
	}
	if tab.Emeritus != nil {
		t.Errorf("expected nil emeritus, got %v", tab.Emeritus)
	}
	if len(config.UserGroups) != 0 {
		t.Errorf("expected 0 user groups, got %d", len(config.UserGroups))
	}
}

// ---------------------------------------------------------------------------
// writeTemplate
// ---------------------------------------------------------------------------

func TestWriteTemplate(t *testing.T) {
	tmpl := template.Must(template.New("test").Parse("Hello {{ .Name }}!"))
	dir := t.TempDir()
	out := filepath.Join(dir, "out.txt")

	data := struct{ Name string }{"World"}
	if err := writeTemplate(tmpl, data, out); err != nil {
		t.Fatalf("writeTemplate failed: %v", err)
	}

	got, _ := os.ReadFile(out)
	want := "Hello World!" + beginCustomMarkdown + "\n" + endCustomMarkdown + "\n"
	if string(got) != want {
		t.Errorf("got %q, want %q", string(got), want)
	}
}

func TestWriteTemplateError(t *testing.T) {
	tmpl := template.Must(template.New("bad").Parse("{{ .Missing.Field }}"))
	dir := t.TempDir()
	out := filepath.Join(dir, "out.txt")

	err := writeTemplate(tmpl, struct{}{}, out)
	if err == nil {
		t.Fatal("expected error from template execution")
	}
}

// ---------------------------------------------------------------------------
// processGroups (integration-style)
// ---------------------------------------------------------------------------

func newTestTemplate(t *testing.T) *template.Template {
	t.Helper()
	funcMap := template.FuncMap{
		"lower":     strings.ToLower,
		"replace":   strings.ReplaceAll,
		"trimSpace": strings.TrimSpace,
		"githubLink": func(p Person) string {
			if p.GitHub != "" {
				return "[" + p.Name + "](https://github.com/" + p.GitHub + ")"
			}
			return p.Name
		},
	}

	groupTmplStr := `# {{ .Name }}
{{ if .Leadership }}
## Leadership
{{ range .Leadership -}}
* {{ githubLink . }} ({{ .Company }})
{{ end -}}
{{ end }}
{{ if .Members }}
## Members
{{ range .Members -}}
* {{ githubLink . }}
{{ end -}}
{{ end }}
{{ if .Emeritus }}
## Emeritus
{{ range .Emeritus -}}
* {{ githubLink . }}
{{ end -}}
{{ end }}
`

	return template.Must(template.New("group").Funcs(funcMap).Parse(groupTmplStr))
}

func TestProcessGroups_Basic(t *testing.T) {
	dir := t.TempDir()
	groupTmpl := newTestTemplate(t)

	groups := []Group{
		{
			Name: "My Group",
			Dir:  "my-group",
			Leadership: []Person{
				{Name: "Alice", GitHub: "alice", Company: "ACME", Role: "Chair"},
			},
		},
	}

	processGroups(groups, "tab", dir, "https://github.com/cncf/tab", groupTmpl)

	readme := filepath.Join(dir, "my-group", "README.md")
	data, err := os.ReadFile(readme)
	if err != nil {
		t.Fatalf("README not created: %v", err)
	}
	content := string(data)
	if !strings.Contains(content, "# My Group") {
		t.Error("missing group name in output")
	}
	if !strings.Contains(content, "[Alice](https://github.com/alice)") {
		t.Error("missing leader link in output")
	}
	if !strings.Contains(content, beginCustomMarkdown) {
		t.Error("missing custom content markers")
	}
}

func TestProcessGroups_WithMembers(t *testing.T) {
	dir := t.TempDir()
	groupTmpl := newTestTemplate(t)

	groups := []Group{
		{
			Name: "Full Group",
			Dir:  "full-group",
			Leadership: []Person{
				{Name: "Lead", GitHub: "lead", Company: "LeadCo", Role: "Chair"},
			},
			Members: []Person{
				{Name: "MemberA", GitHub: "membera", Company: "CorpA"},
				{Name: "MemberB", GitHub: "memberb", Company: "CorpB"},
			},
			Emeritus: []Person{
				{Name: "OldTimer", GitHub: "oldtimer", Company: "PastCo"},
			},
		},
	}

	processGroups(groups, "tab", dir, "", groupTmpl)

	data, err := os.ReadFile(filepath.Join(dir, "full-group", "README.md"))
	if err != nil {
		t.Fatalf("README not created: %v", err)
	}
	content := string(data)

	checks := []string{
		"[Lead](https://github.com/lead)",
		"[MemberA](https://github.com/membera)",
		"[MemberB](https://github.com/memberb)",
		"[OldTimer](https://github.com/oldtimer)",
		"## Leadership",
		"## Members",
		"## Emeritus",
	}
	for _, c := range checks {
		if !strings.Contains(content, c) {
			t.Errorf("output missing %q", c)
		}
	}
}

func TestProcessGroups_AutoDir(t *testing.T) {
	dir := t.TempDir()
	groupTmpl := newTestTemplate(t)

	groups := []Group{
		{
			Name: "Auto Dir Group",
		},
	}

	processGroups(groups, "user-group", dir, "", groupTmpl)

	readme := filepath.Join(dir, "auto-dir-group", "README.md")
	if _, err := os.Stat(readme); err != nil {
		t.Fatalf("auto-dir README not created: %v", err)
	}
}

func TestProcessGroups_PreservesExistingCustomContent(t *testing.T) {
	dir := t.TempDir()
	groupTmpl := newTestTemplate(t)

	// Pre-populate a README with custom content between markers.
	groupDir := filepath.Join(dir, "my-group")
	os.MkdirAll(groupDir, 0o755)
	existing := "# Old\n" + beginCustomMarkdown + "\nMy special notes\n" + endCustomMarkdown + "\n"
	os.WriteFile(filepath.Join(groupDir, "README.md"), []byte(existing), 0o644)

	groups := []Group{
		{
			Name: "My Group",
			Dir:  "my-group",
		},
	}

	processGroups(groups, "tab", dir, "", groupTmpl)

	data, _ := os.ReadFile(filepath.Join(groupDir, "README.md"))
	content := string(data)
	if !strings.Contains(content, "My special notes") {
		t.Error("custom content should be preserved across regeneration")
	}
	if !strings.Contains(content, "# My Group") {
		t.Error("group name should be updated")
	}
	if !strings.Contains(content, beginCustomMarkdown) {
		t.Error("custom content markers should still be present")
	}
}

func TestGetExistingCustomContent_InvalidMarkers(t *testing.T) {
	tests := []struct {
		name    string
		content string
	}{
		{"missing end", beginCustomMarkdown},
		{"missing begin", endCustomMarkdown},
		{"end before begin", endCustomMarkdown + beginCustomMarkdown},
		{"duplicate pair", beginCustomMarkdown + endCustomMarkdown + beginCustomMarkdown + endCustomMarkdown},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path := filepath.Join(t.TempDir(), "README.md")
			if err := os.WriteFile(path, []byte(tt.content), 0o644); err != nil {
				t.Fatal(err)
			}
			if _, err := getExistingCustomContent(path); err == nil {
				t.Fatal("expected invalid marker error")
			}
		})
	}
}

func TestWriteTemplate_PreservesCustomContentVerbatim(t *testing.T) {
	path := filepath.Join(t.TempDir(), "README.md")
	customContent := "\n## Maintainer notes\n\nKeep **this formatting** exactly.\n"
	existing := "Generated content\n" + beginCustomMarkdown + customContent + endCustomMarkdown + "\n"
	if err := os.WriteFile(path, []byte(existing), 0o644); err != nil {
		t.Fatal(err)
	}

	tmpl := template.Must(template.New("test").Parse("Generated content\n"))
	if err := writeTemplate(tmpl, struct{}{}, path); err != nil {
		t.Fatalf("writeTemplate failed: %v", err)
	}

	got, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	want := "Generated content\n" + beginCustomMarkdown + customContent + endCustomMarkdown + "\n"
	if string(got) != want {
		t.Errorf("generated README = %q, want %q", got, want)
	}
}

func TestProcessGroups_MultipleGroups(t *testing.T) {
	dir := t.TempDir()
	groupTmpl := newTestTemplate(t)

	groups := []Group{
		{Name: "Group A", Dir: "group-a"},
		{Name: "Group B", Dir: "group-b"},
		{Name: "Group C"}, // auto-dir
	}

	processGroups(groups, "user-group", dir, "", groupTmpl)

	for _, d := range []string{"group-a", "group-b", "group-c"} {
		readme := filepath.Join(dir, d, "README.md")
		if _, err := os.Stat(readme); err != nil {
			t.Errorf("README not created for %s: %v", d, err)
		}
	}
}

// ---------------------------------------------------------------------------
// Template rendering with real template file
// ---------------------------------------------------------------------------

func TestRealTemplate(t *testing.T) {
	if _, err := os.Stat("group_readme.tmpl"); err != nil {
		t.Skipf("skipping: template not found (run tests from generator/)")
	}

	funcMap := template.FuncMap{
		"lower":     strings.ToLower,
		"replace":   strings.ReplaceAll,
		"trimSpace": strings.TrimSpace,
		"githubLink": func(p Person) string {
			if p.GitHub != "" {
				return "[" + p.Name + "](https://github.com/" + p.GitHub + ")"
			}
			return p.Name
		},
	}

	tmpl := template.Must(template.New("group_readme.tmpl").Funcs(funcMap).ParseFiles("group_readme.tmpl"))

	td := TemplateData{
		Group: Group{
			Name:             "Test Group",
			MissionStatement: "We test things.",
			Leadership: []Person{
				{Name: "Alice", GitHub: "alice", Company: "TestCo", Role: "Chair"},
				{Name: "ViceAlice", GitHub: "vicealice", Company: "TestCo", Role: "Vice Chair"},
			},
			Members: []Person{
				{Name: "MemberX", GitHub: "memberx", Company: "CorpX"},
			},
			Emeritus: []Person{
				{Name: "OldBob", GitHub: "oldbob", Company: "PastCo"},
			},
			Contact: Contact{
				Slack:        "https://slack.example.com/channel",
				SlackChannel: "test-group",
				MailingList:  "test@example.com",
				GitHubTeams:  []GitHubTeam{{Name: "test-team", Description: "Test team"}},
				Liaison:      []Person{{Name: "Liam", GitHub: "liam"}},
			},
			Meetings: []Meeting{
				{
					Description:     "Weekly sync",
					MeetingURL:      "https://example.com/meet",
					MeetingNotesURL: "https://example.com/notes",
					RecordingsURL:   "https://example.com/rec",
				},
			},
		},
		GroupType:   "tab",
		RepoBaseURL: "https://github.com/cncf/tab",
	}

	dir := t.TempDir()
	out := filepath.Join(dir, "README.md")
	if err := writeTemplate(tmpl, td, out); err != nil {
		t.Fatalf("template execution failed: %v", err)
	}

	data, _ := os.ReadFile(out)
	content := string(data)

	checks := []string{
		"# Test Group",
		"We test things.",
		"[Alice](https://github.com/alice)",
		"@alice",
		"TestCo",
		"[ViceAlice](https://github.com/vicealice)",
		"[MemberX](https://github.com/memberx)",
		"[OldBob](https://github.com/oldbob)",
		"#test-group",
		"test@example.com",
		"@cncf/test-team",
		"[Liam](https://github.com/liam)",
		"Weekly sync",
		"https://example.com/meet",
		"Meeting Notes",
		"Recordings",
		beginCustomMarkdown,
		endCustomMarkdown,
	}
	for _, c := range checks {
		if !strings.Contains(content, c) {
			t.Errorf("output missing %q", c)
		}
	}
}

func TestRealTemplate_MinimalGroup(t *testing.T) {
	if _, err := os.Stat("group_readme.tmpl"); err != nil {
		t.Skipf("skipping: template not found (run tests from generator/)")
	}

	funcMap := template.FuncMap{
		"lower":     strings.ToLower,
		"replace":   strings.ReplaceAll,
		"trimSpace": strings.TrimSpace,
		"githubLink": func(p Person) string {
			if p.GitHub != "" {
				return "[" + p.Name + "](https://github.com/" + p.GitHub + ")"
			}
			return p.Name
		},
	}

	tmpl := template.Must(template.New("group_readme.tmpl").Funcs(funcMap).ParseFiles("group_readme.tmpl"))

	td := TemplateData{
		Group: Group{
			Name: "Minimal Group",
		},
		GroupType: "user-group",
	}

	dir := t.TempDir()
	out := filepath.Join(dir, "README.md")
	if err := writeTemplate(tmpl, td, out); err != nil {
		t.Fatalf("template execution failed: %v", err)
	}

	data, _ := os.ReadFile(out)
	content := string(data)

	if !strings.Contains(content, "# Minimal Group") {
		t.Error("missing group name")
	}
	if !strings.Contains(content, beginCustomMarkdown) {
		t.Error("missing custom content markers")
	}
	// Should NOT contain sections for empty fields
	if strings.Contains(content, "## Leadership") {
		t.Error("should not render Leadership section when empty")
	}
	if strings.Contains(content, "## Members") {
		t.Error("should not render Members section when empty")
	}
	if strings.Contains(content, "## Emeritus") {
		t.Error("should not render Emeritus section when empty")
	}
}
