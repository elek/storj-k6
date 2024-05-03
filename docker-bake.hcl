target "build" {
  target = "build"
}

target "image" {
  target = "image"
  output = [ "type=image,name=ghcr.io/elek/storj-k6" ]
}
