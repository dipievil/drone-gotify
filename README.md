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
      title: "Drone Build #{{build.number}}"
      message: "The build {{build.status}} on branch {{commit.branch}}."
      markdown: true      
```

## Environment Variables

When running locally, you can use `.env` to configure the settings:

```dotenv
PLUGIN_URL=https://my-gotify-instance:8081
PLUGIN_TOKEN=my-app-token
PLUGIN_TITLE="Build {{build.number}}"
PLUGIN_MESSAGE="Status: {{build.status}}"
PLUGIN_MARKDOWN=true
```

## License

Licensed under the Apache License 2.0. See [LICENSE](LICENSE).
