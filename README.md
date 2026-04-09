# drone-gotify

Drone plugin to send notifications to [Gotify](https://gotify.net/).

## Usage

Set up the plugin in your `.drone.yml`:

```yaml
steps:
  - name: send-gotify-notification
    image: dipievil/drone-gotify
    settings:
      url: https://my-gotify-instance:8081
      token: your-gotify-app-token
```

If the pipeline failed, the notification will be:

```
❌ Build #42 of `my-repo` failed.
Commit by `john.doe` on `main`:
  chore: update README.md

🌐 https://drone.example.com/my-repo/42
```

If the pipeline succeeded, the notification will be:

```markdown


✅ Build #42 of `my-repo` succeeded.
Commit by `john.doe` on `main`:

chore: update README.md

🌐 https://drone.example.com/my-repo/42
```

More examples:

- **Customize title:**

```yaml
steps:
  - name: send-gotify-notification
    image: dipievil/drone-gotify
    settings:
      url: https://my-gotify-instance:8081
      token: your-gotify-app-token
      title: "Build #{{build.number}} with {{build.status}}"
```

Message:

```markdown
✅ Build #42 of my-repo succeeded.
Commit by john.doe on main:

chore: update README.md

🌐 https://drone.example.com/my-repo/42
```

- **Set high priority:**

```yaml
steps:
  - name: send-gotify-notification
    image: dipievil/drone-gotify
    settings:
      url: https://my-gotify-instance:8081
      token: your-gotify-app-token
      priority: 10
```

Message:

```markdown
✅ Build #42 of `my-repo` succeeded.
Commit by `john.doe` on `main`:

chore: update README.md

🌐 https://drone.example.com/my-repo/42

🟡 Critical priority
```

- **Render message as Markdown:**

```yaml
steps:
  - name: send-gotify-notification
    image: dipievil/drone-gotify
    settings:
      url: https://my-gotify-instance:8081
      token: your-gotify-app-token
      markdown: true
```

Message:

```markdown
✅ Build #42 of `my-repo` succeeded.
Commit by `john.doe` on `main`:

chore: update README.md

🌐 [https://drone.example.com/my-repo/42](https://drone.example.com/my-repo/42)
```

## Environment Variables

| Variable | Required | Default | Description |
| --- | --- | --- | --- |
| `PLUGIN_URL` | Yes | — | Gotify server URL |
| `PLUGIN_TOKEN` | Yes | — | Gotify app token |
| `PLUGIN_TITLE` | No | `Drone CI` | Notification title (supports template variables) |
| `PLUGIN_MESSAGE` | No | — | Notification body (supports template variables) |
| `PLUGIN_PRIORITY` | No | `5` | Message priority |
| `PLUGIN_MARKDOWN` | No | `false` | Render message as Markdown in Gotify clients |
| `PLUGIN_CLICK_URL` | No | — | URL to open when the notification is clicked |
| `PLUGIN_ENV_FILE` | No | `.env` | Path to a custom env file to load at startup |

When running locally, you can use `.env` to configure the settings:

```dotenv
PLUGIN_URL=https://my-gotify-instance:8081
PLUGIN_TOKEN=my-app-token
PLUGIN_TITLE="Build {{build.number}}"
PLUGIN_MESSAGE="Status: {{build.status}}"
PLUGIN_PRIORITY=5
PLUGIN_MARKDOWN=true
PLUGIN_CLICK_URL=https://drone.example.com/{{repo.owner}}/{{repo.name}}/{{build.number}}
```

## License

Licensed under the Apache License 2.0. See [LICENSE](LICENSE).
