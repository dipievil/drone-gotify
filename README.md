# drone-gotify

Drone plugin to send notifications to [Gotify](https://gotify.net/). 

## Usage

Set up the plugin in your `.drone.yml`:

```yaml
steps:
  - name: send-gotify-notification
    image: plugins/gotify
    settings:
      url: https://gotify.example.com
      token: your-gotify-app-token
      title: "Drone Build #{{build.number}}"
      message: "The build {{build.status}} on branch {{commit.branch}}."
      markdown: true
      click_url: "{{build.link}}"
```

## Environment Variables

When running locally, you can use `.env` to configure the settings:

```dotenv
PLUGIN_URL=https://gotify.example.com
PLUGIN_TOKEN=my-app-token
PLUGIN_TITLE="Build {{build.number}}"
PLUGIN_MESSAGE="Status: {{build.status}}"
PLUGIN_MARKDOWN=true
```
