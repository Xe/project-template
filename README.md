# project-template

A template with yeet, semantic-release, and more set up.

## Project structure

- `cmd`: Executable commands / entrypoints
- `cmd/web`: The HTTP server for this app
- `cmd/web/server.go`: Where all your server logic should live as methods on Server
- `globals`: Global constants such as the version number (updated automatically at build time)
- `internal`: Where harsh realities that shouldn't or can't be reused should go
- `models`: Where database models should go. Make your actions methods on the DAO struct [like this](https://github.com/Xe/x/blob/master/cmd/mi/models/blogpost.go)
- `web`: Where HTML templates for this app should go using [templ](https://templ.guide/). See the templ docs for more information and also look through `web/index.templ` for ideas. Make one template file per area of concern.
- `web/static`: Where static files should go. These files will be compiled into the binary and served at `/static/...`.

## Usage

1. Hit "Use this template" on Gitea and create a new repo
1. Choose a user / org for it (either your username or Techaro)
1. Choose the following template items:

   - Git Content (Default Branch)
   - Issue Labels

1. Hit "Create Repository"
1. (if you aren't making this in the Techaro org) Open the repo settings, click on collaboration, invite [Mimi](https://git.xeserv.us/mimi) with Administrator permissions so she can make releases for you
1. Clone to your machine
1. Open in a development container
1. Edit package.json:

   ```diff
   -"name": "@xe/project-template",
   +"name": "@techaro/project-name",
   -"version": "<whatever>",
   +"version": "0.0.1",
   ```

1. Rename the package based on your git repo URL:

   ```text
   npm run rename git.xeserv.us/Techaro/project-name
   ```

1. Commit the new data to main:

   ```text
   git add .
   git commit -sm "chore: initial commit"
   ```

1. Push to Gitea:

   ```text
   git push
   ```

1. Open the Actions view in the repo on gitea and be sure tests pass

## Local tasks

To forcibly regenerate generated files:

```text
npm run generate
```

To spawn an instance of this in development:

```text
npm run dev
```

To open a shell to the database:

```text
npm run dev:psql
```

To open a shell to Redis/Valkey:

```text
npm run dev:redis
```

To format your code locally:

```text
npm run format
```

To run tests locally:

```text
npm run test
```

To run what CI runs locally:

```text
npm run test:gha
```
