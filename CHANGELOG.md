# [0.9.0](http://git.gitea.svc.alrest.xeserv.us/xe/project-template/compare/v0.8.0...v0.9.0) (2025-07-19)

### Bug Fixes

- **dao:** add advisory lock to avoid race conditions when multiple tests migrate at the same time ([d929ea0](http://git.gitea.svc.alrest.xeserv.us/xe/project-template/commit/d929ea068de52a2bfa76f6180f92c834a22999de))

### Features

- add valkey gorm caching, sessions in valkey ([acf715e](http://git.gitea.svc.alrest.xeserv.us/xe/project-template/commit/acf715ec3fee55ce62409ec05dd5d31009fa57d1))
- implement session middleware and flash message storage ([4240007](http://git.gitea.svc.alrest.xeserv.us/xe/project-template/commit/4240007d80408a1c4c6b514022139edf0707c6ba))
- **web:** compress static assets on the wire, import htmx ([46bf4a7](http://git.gitea.svc.alrest.xeserv.us/xe/project-template/commit/46bf4a778611a892267cb8bab8b523d6aac7ee81))

# [0.8.0](http://git.gitea.svc.alrest.xeserv.us/xe/project-template/compare/v0.7.0...v0.8.0) (2025-07-15)

### Features

- **devcontainer/personalize:** rewrite to totally handle all of the previously manual setup logic ([8f42d5c](http://git.gitea.svc.alrest.xeserv.us/xe/project-template/commit/8f42d5cc453b77fbcf2a60dd2ed7d06c4c357117))

# [0.7.0](http://git.gitea.svc.alrest.xeserv.us/xe/project-template/compare/v0.6.0...v0.7.0) (2025-07-15)

### Features

- workspace folder personalization script ([a2e651d](http://git.gitea.svc.alrest.xeserv.us/xe/project-template/commit/a2e651d48de9e0054d48ae8c170935a641d4f853))

# [0.6.0](http://git.gitea.svc.alrest.xeserv.us/xe/project-template/compare/v0.5.0...v0.6.0) (2025-07-15)

### Bug Fixes

- move xess into web ([89fd955](http://git.gitea.svc.alrest.xeserv.us/xe/project-template/commit/89fd95555a5b1365d8763131438d3981be2a4712))

### Features

- add ability to auto rename all code ([c287808](http://git.gitea.svc.alrest.xeserv.us/xe/project-template/commit/c2878087ecf8ebbce6462c91833692f0c4a63673))

# [0.5.0](https://git.gitea.svc.alrest.xeserv.us/xe/project-template/compare/v0.4.0...v0.5.0) (2025-07-15)

### Features

- add dao models ([a936c7e](https://git.gitea.svc.alrest.xeserv.us/xe/project-template/commit/a936c7e2f46ef3947de5d86eb092c40cac87188b))
- add docker metadata ([d693df3](https://git.gitea.svc.alrest.xeserv.us/xe/project-template/commit/d693df33570c6ea93caae3569d316c11bf94d16f))
- add logging scaffolding ([24a5048](https://git.gitea.svc.alrest.xeserv.us/xe/project-template/commit/24a5048d99a3bbe2933d9c3838449ecc18681fcf))
- add xess and basic HTML templates ([13d5241](https://git.gitea.svc.alrest.xeserv.us/xe/project-template/commit/13d5241a76893faf13b33c1e45ddbb4c4f283052))
- format on generate ([564eb3a](https://git.gitea.svc.alrest.xeserv.us/xe/project-template/commit/564eb3a211d595dcd4208a2e88a5bd8056a0dc5c))
- move main.go to ./cmd/web, add basic setup ([a79a012](https://git.gitea.svc.alrest.xeserv.us/xe/project-template/commit/a79a012050aa4ff36340e7f796a4420777599765))
- move server logic to cmd/web/server.go ([839bb4b](https://git.gitea.svc.alrest.xeserv.us/xe/project-template/commit/839bb4bc427a0630764c39a741c98bb773a903d3))
- npm run format ([b572e06](https://git.gitea.svc.alrest.xeserv.us/xe/project-template/commit/b572e069eadaf3187d7de63dcd4b7d26a4fe7de9))

# [0.4.0](https://git.gitea.svc.alrest.xeserv.us/xe/project-template/compare/v0.3.0...v0.4.0) (2025-07-04)

### Features

- **devcontainer:** use docker compose ([b0bfe42](https://git.gitea.svc.alrest.xeserv.us/xe/project-template/commit/b0bfe42fe9f3dee84111f99d779d9cfa3189cf05))

# [0.3.0](https://git.gitea.svc.alrest.xeserv.us/xe/project-template/compare/v0.2.5...v0.3.0) (2025-07-03)

### Features

- development container ([3bcb676](https://git.gitea.svc.alrest.xeserv.us/xe/project-template/commit/3bcb67647764c7b0a33ee8647b55f2aa55bab8d9))

## [0.2.5](https://git.gitea.svc.alrest.xeserv.us/xe/project-template/compare/v0.2.4...v0.2.5) (2025-05-03)

### Bug Fixes

- **release:** do this the less braindead way ([f3930f7](https://git.gitea.svc.alrest.xeserv.us/xe/project-template/commit/f3930f757f9ae5dca7415c4e2266f70b9bec4be1))

## [0.2.4](https://git.gitea.svc.alrest.xeserv.us/xe/project-template/compare/v0.2.3...v0.2.4) (2025-05-03)

### Bug Fixes

- **release:** git cat-file -p HEAD ([ea6d110](https://git.gitea.svc.alrest.xeserv.us/xe/project-template/commit/ea6d11071a30a53d5ac98ab06684e5cee5a4eb59))

## [0.2.3](https://git.gitea.svc.alrest.xeserv.us/xe/project-template/compare/v0.2.2...v0.2.3) (2025-05-03)

### Bug Fixes

- **release:** do you do --debug this way? ([4525263](https://git.gitea.svc.alrest.xeserv.us/xe/project-template/commit/4525263603aa3032f67f2e7f2ce94623d576c51c))

## [0.2.2](https://git.gitea.svc.alrest.xeserv.us/xe/project-template/compare/v0.2.1...v0.2.2) (2025-05-03)

### Bug Fixes

- **ci:** give our robot overlord her own key ([49a4a13](https://git.gitea.svc.alrest.xeserv.us/xe/project-template/commit/49a4a13dcaac760c81af6706d553cb2f09a81b1d))

## [0.2.1](https://git.gitea.svc.alrest.xeserv.us/xe/project-template/compare/v0.2.0...v0.2.1) (2025-05-03)

### Bug Fixes

- re-sign Mimi's keys ([de6460c](https://git.gitea.svc.alrest.xeserv.us/xe/project-template/commit/de6460c0058d3151c94ccf00c6ef943f830b4a14))

# [0.2.0](https://git.gitea.svc.alrest.xeserv.us/xe/project-template/compare/v0.1.0...v0.2.0) (2025-05-03)

### Features

- sign commits with GPG ([163b553](https://git.gitea.svc.alrest.xeserv.us/xe/project-template/commit/163b553eb6634acc6604f4e1cbdf763815f365a1))

# [0.2.0](https://git.gitea.svc.alrest.xeserv.us/xe/project-template/compare/v0.1.0...v0.2.0) (2025-05-03)

### Features

- sign commits with GPG ([163b553](https://git.gitea.svc.alrest.xeserv.us/xe/project-template/commit/163b553eb6634acc6604f4e1cbdf763815f365a1))

# [0.1.0](https://git.gitea.svc.alrest.xeserv.us/xe/project-template/compare/v0.0.2...v0.1.0) (2025-05-03)

### Features

- **release:** give Mimi commit credit ([d3b6163](https://git.gitea.svc.alrest.xeserv.us/xe/project-template/commit/d3b61637dadc7503ae4b49b96971f97103bb4146))

## [0.0.2](https://git.gitea.svc.alrest.xeserv.us/xe/project-template/compare/v0.0.1...v0.0.2) (2025-05-03)
