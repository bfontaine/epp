#! /bin/bash
for os in linux darwin windows; do
  for arch in amd64 386; do
    if [[ "$os-$arch" = "darwin-386" ]]; then continue; fi
    name="epp-$os-$arch"
    # -s: Omit the symbol table and debug information.
    # -w: Omit the DWARF symbol table.
    env GOOS=$os GOARCH=$arch go build -ldflags="-s -w" -o $name .
    shasum -a 256 $name
  done
done
