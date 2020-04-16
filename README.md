Project Manager Web Site
==================

# UNDER DEVELOPMENT

TODO: more features, docs, multi-user environment

Uses awesome [mithril.js](https://mithril.js.org/) framework as front-end, [golang](https://golang.org/) as back-end service.

Supports time tracking, markdown syntax & file attachments

Prerequisites: go, node.js with npm, postgresql server installed

Call `npm i` to install client npm libraries, `go get ./...` to install all golang dependencies. Copy `config/config.yml.example` to `config/config.yml` and adjust settings as needed.

Call `make watch` to build & watch assets, `fresh` to run & watch golang backend api-service during development.