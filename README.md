# project-template

A template with Docker, xess, Postgres, semantic-release, and more set up.

Things you need to know:

- This template enforces [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/). They also ship a [cheat sheet](https://www.conventionalcommits.org/en/v1.0.0/#summary).
- This template works best with [Development containers](https://containers.dev/) in your editor.

## Services

When you start this project as a development container, it spawns the following services:

- [Postgres](https://www.postgresql.org/) version 16 for permanent data storage.
- [Valkey](https://valkey.io/) version 8 for cache data or ephemeral data stroage.

Prefer using Postgres for storing data.

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
   -"version": "<whatever>",
   +"name": "@techaro/project-name",
   +"version": "0.0.1",
   ```

1. Remove `CHANGELOG.md`:

   ```text
   rm CHANGELOG.md
   ```

1. Rename the package based on your git repo URL:

   ```text
   npm run rename git.xeserv.us/Techaro/project-name
   ```

1. Edit `docker-bake.hcl` based on your project name:

   ```diff
      tags = [
   -    "registry.int.xeserv.us/projects/name:latest",
   -    "registry.int.xeserv.us/projects/name:${GITHUB_SHA}"
   +    "registry.int.xeserv.us/owner/name:latest",
   +    "registry.int.xeserv.us/owner/name:${GITHUB_SHA}"
      ]
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
1. If tests pass, tag version `v0.0.1`:

   ```text
   git tag -s v0.0.1 -m "initial tag"
   ```

1. Then push to Gitea:

   ```text
   git push
   git push --tags
   ```

1. Gitea will process things and then push version v0.1.0. This is normal. Once it's done, `git pull`:

   ```text
   git pull
   ```

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
