variable "ALPINE_VERSION" { default = "edge" }
variable "GO_VERSION" { default = "1.24" }
variable "GITHUB_SHA" { default = "devel" }

group "default" {
  targets = [
    "web",
  ]
}

target "web" {
  args = {
    ALPINE_VERSION = null
    GO_VERSION = null
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