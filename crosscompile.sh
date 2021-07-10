#! /bin/bash
for os in linux darwin windows; do
  for arch in amd64 386; do
    if [[ "$os-$arch" = "darwin-386" ]]; then continue; fi
    name="epp-$os-$arch"
    env GOOS=$os GOARCH=$arch go build -o $name .
    shasum -a 256 $name
  done
done
