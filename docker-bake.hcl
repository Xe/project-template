variable "ALPINE_VERSION" { default = "3.22" }
variable "GO_VERSION" { default = "1.24" }
variable "GITHUB_SHA" { default = "devel" }

group "default" {
  targets = [
    "web",
  ]
}

target "devcontainer" {
  context = "."
  dockerfile = ".devcontainer/Dockerfile"
  platforms = [
    "linux/amd64",
    "linux/arm64",
  ]
  cache-from = [
    "type=registry,ref=ghcr.io/xe/project-template/devcontainer/cache"
  ]
  cache-to = [
    "type=registry,ref=ghcr.io/xe/project-template/devcontainer/cache"
  ]
  pull = true
  tags = [
    "ghcr.io/xe/project-template/devcontainer:latest"
  ]
}

target "web" {
  args = {
    ALPINE_VERSION = "${ALPINE_VERSION}"
    GO_VERSION = "${GO_VERSION}"
  }
  context = "."
  dockerfile = "./Dockerfile"
  platforms = [
    "linux/amd64",
    "linux/arm64",
  ]
  pull = true
  tags = [
    "ghcr.io/xe/project-template:latest",
    "ghcr.io/xe/project-template:${GITHUB_SHA}"
  ]
}