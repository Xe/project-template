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

1. Hit "Use this template" on Gitea and create a new repo.
1. Choose a user / org for it (either your username or Techaro).
1. Choose the following template items:

   - Git Content (Default Branch)
   - Issue Labels

1. Hit "Create Repository".
1. (if you aren't making this in the Techaro org) Open the repo settings, click on collaboration, invite [Mimi](https://github.com/Xe) with Administrator permissions so she can make releases for you.
1. Clone to your machine.
1. Run `.devcontainer/personalize.sh`:

   ```text
   bash .devcontainer/personalize.sh
   ```

   This does the following:

   - Customizes the mountpount of the development container workspace folder so it matches your project name.
   - Edits the development container docker compose to match that workspace folder change.
   - Renames all of the Go imports to your Gitea repo so you don't try to pull code from the template on accident.
   - Updates the `name` in `package.json` to `@you/project`.
   - Updates the `version` in `package.json` to `0.0.0` (Mimi will manage versioning for you).
   - Removes some setup logic from the `package.json` `scripts` section.
   - Deletes the [CHANGELOG.md](./CHANGELOG.md) file that the template uses.
   - Edits `docker-bake.hcl` to make a unique image for you.
   - Cleans up other tooling dependencies for setup.
   - Formats your code.
   - Deletes the setup script.

1. Open in a development container.
1. Commit the new data to main:

   ```text
   git add .
   git commit -sm "feat: initial commit"
   ```

1. Push to Gitea:

   ```text
   git push
   ```

1. Open the Actions view in the repo on gitea and be sure tests pass.
1. If tests pass, open the Actions view, click on `release.yaml`, and then click "Run Workflow". You may need to run it twice.
1. Gitea will process things and then Mimi will push version v1.0.0. This is normal. Once it's done, `git pull`:

   ```text
   git pull
   ```

You can now develop your service as normal!

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
