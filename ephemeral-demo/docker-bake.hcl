group "default" {
  targets = ["app"]
}

variable "REGISTRY" { default = "local" }
variable "IMAGE"    { default = "preview-app" }
variable "TAG"      { default = "demo" }
variable "VERSION"  { default = "dev" }
variable "COMMIT"   { default = "" }

function "get_commit" {
  params = []
  result = notequal("", COMMIT) ? COMMIT : ""
}

function "get_build_time" {
  params = []
  result = timestamp()
}

target "app" {
  context = "app"
  dockerfile = "Dockerfile"
  tags = ["${REGISTRY}/${IMAGE}:${TAG}"]
  args = {
    VERSION = VERSION
    COMMIT = get_commit()
    BUILD_TIME = get_build_time()
  }
  cache-to = ["type=inline"]
  sbom = true
  provenance = true
}
