target "hello-world-app" {
  context = "."
  dockerfile = "Dockerfile"
  args = {
    HELLO_MSG = "Hello from Docker Bake!"
  }
  tags = ["hello-world-app:latest"]
  platforms = ["linux/amd64", "linux/arm64"]
}
