package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/appleboy/drone-template-lib/template"
)

type (
	// GitHub information.
	GitHub struct {
		Workflow  string
		Workspace string
		Action    string
		EventName string
		EventPath string
	}

	// Repo information.
	Repo struct {
		FullName  string
		Namespace string
		Name      string
	}

	// Commit information.
	Commit struct {
		Sha     string
		Ref     string
		Branch  string
		Link    string
		Author  string
		Avatar  string
		Email   string
		Message string
	}

	// Build information.
	Build struct {
		Tag      string
		Event    string
		Number   int
		Status   string
		Link     string
		Started  int64
		Finished int64
		PR       string
		DeployTo string
	}

	// Config for the plugin.
	Config struct {
		URL      string
		Token    string
		Title    string
		Message  string
		Priority int
		Markdown bool
		ClickURL string
		GitHub   bool
	}

	// Plugin values.
	Plugin struct {
		GitHub GitHub
		Repo   Repo
		Commit Commit
		Build  Build
		Config Config
		Tpl    map[string]string
	}

	GotifyMessage struct {
		Title    string                 `json:"title"`
		Message  string                 `json:"message"`
		Priority int                    `json:"priority"`
		Extras   map[string]interface{} `json:"extras,omitempty"`
	}
)

var icons = map[string]string{
	"failure":   "❌",
	"cancelled": "❕",
	"success":   "✅",
}

func templateMessage(t string, plugin Plugin) (string, error) {
	return template.RenderTrim(t, plugin)
}

// DefaultMessage returns the plugin default message if not provided.
func (p Plugin) DefaultMessage() string {
	icon := icons[strings.ToLower(p.Build.Status)]

	if p.Config.GitHub {
		return fmt.Sprintf("%s/%s triggered by %s (%s)",
			p.Repo.FullName,
			p.GitHub.Workflow,
			p.Repo.Namespace,
			p.GitHub.EventName,
		)
	}

	return fmt.Sprintf("%s Build #%d of `%s` %s.\n\n📝 Commit by %s on `%s`:\n``` %s ```\n\n🌐 %s",
		icon,
		p.Build.Number,
		p.Repo.FullName,
		p.Build.Status,
		p.Commit.Author,
		p.Commit.Branch,
		p.Commit.Message,
		p.Build.Link,
	)
}

func (p Plugin) buildExtras() map[string]interface{} {
	extras := make(map[string]interface{})

	if p.Config.Markdown {
		extras["client::display"] = map[string]interface{}{
			"contentType": "text/markdown",
		}
	}

	if p.Config.ClickURL != "" {
		extras["client::notification"] = map[string]interface{}{
			"click": map[string]interface{}{
				"url": p.Config.ClickURL,
			},
		}
	}

	if len(extras) == 0 {
		return nil
	}
	return extras
}

// Exec executes the plugin.
func (p Plugin) Exec() error {
	if p.Config.URL == "" {
		return errors.New("missing gotify url")
	}
	if p.Config.Token == "" {
		return errors.New("missing gotify token")
	}

	var message string
	var err error

	if p.Config.Message != "" {
		message, err = templateMessage(p.Config.Message, p)
		if err != nil {
			return fmt.Errorf("error templating message: %w", err)
		}
	} else {
		message = p.DefaultMessage()
	}

	title, err := templateMessage(p.Config.Title, p)
	if err != nil {
		return fmt.Errorf("error templating title: %w", err)
	}

	msg := GotifyMessage{
		Title:    title,
		Message:  message,
		Priority: p.Config.Priority,
		Extras:   p.buildExtras(),
	}

	return p.Send(msg)
}

// Send sends the message to the Gotify server
func (p *Plugin) Send(msg GotifyMessage) error {
	payload, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("error marshaling payload: %w", err)
	}

	parsedURL, err := url.Parse(p.Config.URL)
	if err != nil {
		return fmt.Errorf("error parsing url: %w", err)
	}

	parsedURL.Path = "/message"
	q := parsedURL.Query()
	q.Set("token", p.Config.Token)
	parsedURL.RawQuery = q.Encode()

	req, err := http.NewRequest("POST", parsedURL.String(), bytes.NewBuffer(payload))
	if err != nil {
		return fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("gotify returned status code: %d, response: %s", resp.StatusCode, string(body))
	}

	return nil
}
