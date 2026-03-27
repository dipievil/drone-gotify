---
date: 2026-03-27T00:00:00+00:00
title: Gotify
author: dipi
tags: [ gotify, notifications, chat ]
repo: drone-gotify/drone-gotify
logo: gotify.svg
image: plugins/gotify
---

The Gotify plugin posts build status messages to a Gotify server. The below pipeline configuration demonstrates simple usage:

```yaml
steps:
- name: gotify
  image: plugins/gotify
  settings:
    url: https://gotify.example.com
    token: secret_app_token
    title: "Drone CI Build"
    message: "Build finished!"
```

Example configuration with markdown and click URL:

```yaml
steps:
- name: gotify
  image: plugins/gotify
  settings:
    url: https://gotify.example.com
    token: secret_app_token
    title: "Drone CI: {{repo.name}}"
    message: "Build {{build.number}} triggered by {{commit.author}} {{build.status}}."
    markdown: true
    click_url: "https://drone.example.com/{{repo.owner}}/{{repo.name}}/{{build.number}}"
```

## Parameter Reference

url
: URL for Gotify server

token
: App token to authenticate with

title
: Title for the message. Evaluates template variables.

message
: Body for the message. Evaluates template variables.

priority
: Priority of the message, defaults to 5.

markdown
: Boolean. Set to true to include the `client::display` extra indicating the message is formatted as Markdown text.

click_url
: Provides a `click` url within the `client::notification` extras map.

## Template Variables

The `message` and `title` variables use the Go templating language and can be customized with data from the Drone payload. The following template variables are available to customize your notification:

* `build.status`
* `build.number`
* `build.event`
* `build.link`
* `build.tag`
* `repo.name`
* `repo.namespace`
* `repo.full_name`
* `commit.message`
* `commit.author`
* `commit.sha`
* `commit.ref`
* `commit.branch`
