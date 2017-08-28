#! /bin/bash
for os in linux darwin windows; do
  for arch in amd64 386; do
    name="epp-$os-$arch"
    env GOOS=$os GOARCH=$arch go build -o $name .
    shasum -a 256 $name
  done
done
