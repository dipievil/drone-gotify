# drone-gotify

![Go](https://img.shields.io/badge/Go-1.22-00ADD8?logo=go&logoColor=white) | ![Docker](https://img.shields.io/badge/Docker-Ready-2496ED?logo=docker&logoColor=white) | ![License](https://img.shields.io/github/license/dipievil/drone-gotify)
![Drone CI](https://img.shields.io/badge/Drone%20CI-Compatible-success?logo=drone)

A lightweight Drone CI plugin to send build notifications to [Gotify](https://gotify.net/) with customizable messages and formatting.

## Table of Contents

- [About](#about)
- [Features](#features)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
- [Usage](#usage)
  - [Basic Configuration](#basic-configuration)
  - [Examples](#examples)
  - [Template Variables](#template-variables)
- [Environment Variables](#environment-variables)
- [Configuration](#configuration)
- [Troubleshooting](#troubleshooting)
- [Contributing](#contributing)
- [License](#license)

## About

**drone-gotify** is a plugin for [Drone CI](https://www.drone.io/) that integrates seamlessly with [Gotify](https://gotify.net/), a simple server for sending and receiving notifications. Use this plugin to send build status updates, test results, and custom notifications to your Gotify instance directly from your CI/CD pipeline.

### Features

- **Build Status Notifications**: Automatically send pass/fail status with custom messages
- **Template Support**: Use Drone CI variables in titles and messages (e.g., `{{build.number}}`, `{{repo.name}}`)
- **Markdown Formatting**: Render messages as Markdown in Gotify clients
- **Priority Control**: Set message priority levels (1-10) for important notifications
- **Click Actions**: Add clickable URLs to notifications for quick access to build details
- **Environment-based Configuration**: Configure via environment variables or `.env` files
- **Docker-ready**: Pre-built Docker image available on Docker Hub

## Getting Started

### Prerequisites

- A running Gotify server instance
- A valid Gotify app token
- Drone CI configured in your repository

## Usage

### Basic Configuration

Add the plugin to your `.drone.yml`:

```yaml
steps:
  - name: notify-gotify
    image: dipievil/drone-gotify
    settings:
      url: https://my-gotify-instance:8081
      token: your-gotify-app-token
```

### Examples

#### Example 1: Success/Failure Notifications

By default, the plugin sends build status notifications:

**On success:**
```
✅ Build #42 of `my-repo` succeeded.
Commit by `john.doe` on `main`:

chore: update README.md

🌐 https://drone.example.com/my-repo/42
```

**On failure:**
```
❌ Build #42 of `my-repo` failed.
Commit by `john.doe` on `main`:

chore: update README.md

🌐 https://drone.example.com/my-repo/42
```



#### Example 2: Custom Title

Customize the notification title with template variables:

```yaml
steps:
  - name: notify-gotify
    image: dipievil/drone-gotify
    settings:
      url: https://my-gotify-instance:8081
      token: your-gotify-app-token
      title: "Build #{{build.number}} - {{repo.name}} [{{build.status}}]"
```

#### Example 3: High Priority Notification

Mark important builds as critical:

```yaml
steps:
  - name: notify-gotify
    image: dipievil/drone-gotify
    settings:
      url: https://my-gotify-instance:8081
      token: your-gotify-app-token
      priority: 10
```

Message output:
```
✅ Build #42 of `my-repo` succeeded.
Commit by `john.doe` on `main`:

chore: update README.md

🌐 https://drone.example.com/my-repo/42

🟡 Critical priority
```

#### Example 4: Markdown Formatting

Render message as Markdown in Gotify clients:

```yaml
steps:
  - name: notify-gotify
    image: dipievil/drone-gotify
    settings:
      url: https://my-gotify-instance:8081
      token: your-gotify-app-token
      markdown: true
```

#### Example 5: Custom Click URL

Add a clickable link to notification:

```yaml
steps:
  - name: notify-gotify
    image: dipievil/drone-gotify
    settings:
      url: https://my-gotify-instance:8081
      token: your-gotify-app-token
      click_url: "https://drone.example.com/{{repo.owner}}/{{repo.name}}/{{build.number}}"
```

### Template Variables

The `title` and `message` fields support Go template variables. Available variables include:

**Build Variables:**
- `{{build.status}}` - Build status (success/failure)
- `{{build.number}}` - Build number
- `{{build.event}}` - Build event type
- `{{build.link}}` - Link to the build
- `{{build.tag}}` - Git tag (if applicable)

**Repository Variables:**
- `{{repo.name}}` - Repository name
- `{{repo.namespace}}` - Repository namespace/organization
- `{{repo.full_name}}` - Full repository name (owner/repo)

**Commit Variables:**
- `{{commit.message}}` - Commit message
- `{{commit.author}}` - Commit author
- `{{commit.sha}}` - Commit SHA
- `{{commit.ref}}` - Commit reference
- `{{commit.branch}}` - Branch name

## Environment Variables

Configure the plugin using environment variables. When running in Docker containers, these should be set as `PLUGIN_*` variables:

| Variable | Required | Default | Description |
|----------|----------|---------|-------------|
| `PLUGIN_URL` | ✅ Yes | — | Gotify server URL (e.g., `https://my-gotify-instance:8081`) |
| `PLUGIN_TOKEN` | ✅ Yes | — | Gotify app token for authentication |
| `PLUGIN_TITLE` | No | `Drone CI` | Notification title (supports template variables) |
| `PLUGIN_MESSAGE` | No | — | Notification body (supports template variables) |
| `PLUGIN_PRIORITY` | No | `5` | Message priority (1-10, higher = more critical) |
| `PLUGIN_MARKDOWN` | No | `false` | Render message as Markdown in Gotify clients (`true`/`false`) |
| `PLUGIN_CLICK_URL` | No | — | URL to open when notification is clicked (supports template variables) |
| `PLUGIN_ENV_FILE` | No | `.env` | Path to custom env file to load at startup |

## Configuration

### Using .env File

When running locally, create a `.env` file in the project root:

```dotenv
PLUGIN_URL=https://my-gotify-instance:8081
PLUGIN_TOKEN=my-app-token
PLUGIN_TITLE="Build {{build.number}} - {{repo.name}}"
PLUGIN_MESSAGE="Status: {{build.status}}\nAuthor: {{commit.author}}"
PLUGIN_PRIORITY=5
PLUGIN_MARKDOWN=true
PLUGIN_CLICK_URL=https://drone.example.com/{{repo.owner}}/{{repo.name}}/{{build.number}}
```

### In Docker Compose

```yaml
services:
  drone-gotify:
    image: dipievil/drone-gotify
    environment:
      PLUGIN_URL: https://my-gotify-instance:8081
      PLUGIN_TOKEN: my-app-token
      PLUGIN_PRIORITY: "8"
      PLUGIN_MARKDOWN: "true"
```

## Troubleshooting

### "Connection refused" or "Cannot reach Gotify server"

**Problem:** The plugin cannot connect to your Gotify instance.

**Solutions:**
- Verify the `PLUGIN_URL` is correct and reachable from your Drone runner
- Check that the Gotify server is running: `curl https://my-gotify-instance:8081`
- For local development, ensure the Docker container can access your Gotify instance network
- Check firewall rules and network policies

### "Unauthorized" or "Invalid Token"

**Problem:** Authentication with Gotify failed.

**Solutions:**
- Verify the `PLUGIN_TOKEN` is a valid Gotify app token
- Create a new app token in Gotify admin panel if token has expired
- Ensure the token is not accidentally truncated or modified

### Build succeeds but no notification is sent

**Problem:** Plugin runs without errors but notification doesn't appear.

**Solutions:**
- Check Gotify server logs for received messages
- Verify `PLUGIN_URL` and `PLUGIN_TOKEN` are set correctly
- Enable debug logging if available
- Verify the notification didn't go to a different client/channel

### Template variables not rendering

**Problem:** Variables like `{{build.number}}` appear as literal text.

**Solutions:**
- Ensure you're using double curly braces `{{variable}}`
- Check the variable name matches the list in [Template Variables](#template-variables) section
- Verify template is in `title` or `message` field, not in other settings

### Build on custom network doesn't reach Gotify

**Problem:** Plugin fails to connect when Drone runs on a Docker network.

**Solutions:**
- Use the Gotify container name instead of `localhost`: `http://gotify-container:8080`
- Ensure both containers are on the same Docker network
- Use `host.docker.internal` for host machine access on Docker Desktop

## Development

### Prerequisites for Building

- Go 1.22 or higher
- Make
- Docker (for building container image)

### Building Locally

```bash
# Clone repository
git clone https://github.com/dipievil/drone-gotify.git
cd drone-gotify

# Install dependencies
make tidy

# Build binary
make build

# Run tests
make test

# Run the plugin locally
make run
```

### Running Tests

```bash
go test -v -cover ./...
```

## Contributing

Contributions are welcome! Please follow these guidelines:

1. **Fork the repository** and create a feature branch
   ```bash
   git checkout -b feature/your-feature-name
   ```

2. **Make your changes** and ensure tests pass
   ```bash
   make test
   ```

3. **Follow conventional commits** for commit messages
   - `feat:` for new features
   - `fix:` for bug fixes
   - `docs:` for documentation
   - `test:` for tests
   - `chore:` for maintenance

4. **Create a pull request** with a clear description of your changes

### Code Style

- Follow Go conventions and best practices
- Use `gofmt` for code formatting
- Run tests before submitting a PR
- Update README if adding new features

### Reporting Issues

Please report issues with:
- Clear title describing the problem
- Steps to reproduce
- Expected vs actual behavior
- Environment details (Drone version, Go version, etc.)
- Any error messages or logs

## License

This project is licensed under the **Apache License 2.0**. See [LICENSE](LICENSE) file for details.

## Acknowledgements

- [Gotify](https://gotify.net/) - Simple server for sending and receiving messages
- [Drone CI](https://www.drone.io/) - The leading open-source continuous integration platform
- [urfave/cli](https://github.com/urfave/cli) - A simple, fast, and fun package for building command line apps in Go
- [godotenv](https://github.com/joho/godotenv) - A Go port of the Ruby dotenv project

---

**Questions or Issues?** [Create an issue](https://github.com/dipievil/drone-gotify/issues) on GitHub!
