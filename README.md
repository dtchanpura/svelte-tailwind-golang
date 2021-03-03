# Svelte TailwindCSS and Golang

This is a template repository for Svelte and TailwindCSS with backend as Golang server.

## Features

* Debug flag to allow direct run like `npm run dev` via same port.
* Build part uses `embed` package for production use.

> The inspiration for this template comes from [matryer/crunchcrunchcrunchstack](https://github.com/matryer/crunchcrunchcrunchstack)

## Usage

```shell
# Build and run Go Backend
go run main.go -debug &

# Go to frontend code and start watching
cd frontend

# Optionally do `npm install` if needed
npm run dev

# open http://127.0.0.1:8001 that should open with auto-reload for UI.
```

# LICENSE

MIT