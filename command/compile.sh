#!/bin/bash
go generate

echo "compile each binary"

if [ -z "$TRAVIS_BUILD_DIR" ]
then
	TRAVIS_BUILD_DIR=$PWD
fi

echo $TRAVIS_BUILD_DIR

for GOOS in darwin linux windows; do
  for GOARCH in 386 amd64; do
    echo "Building $GOOS-$GOARCH"
    export GOOS=$GOOS
    export GOARCH=$GOARCH
    export CGO_ENABLED=0
    go build -o $TRAVIS_BUILD_DIR/bin/goNest-$GOOS-$GOARCH -ldflags "-X main.Version=`cat VERSION`"
  done
done
mv bin/goNest-darwin-386 bin/goNest-darwin-386.bin
mv bin/goNest-darwin-amd64 bin/goNest-darwin-amd64.bin
mv bin/goNest-linux-386 bin/goNest-linux-386.bin
mv bin/goNest-linux-amd64 bin/goNest-linux-amd64.bin
mv bin/goNest-windows-386 bin/goNest-windows-386.exe
mv bin/goNest-windows-amd64 bin/goNest-windows-amd64.exe
