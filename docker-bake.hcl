variable "ALPINE_VERSION" { default = "3.22" }
variable "GO_VERSION" { default = "1.24" }
variable "GITHUB_SHA" { default = "devel" }

group "default" {
  targets = [
    "web",
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
    "registry.int.xeserv.us/projects/name:latest",
    "registry.int.xeserv.us/projects/name:${GITHUB_SHA}"
  ]
}