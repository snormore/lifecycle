#/bin/bash
set -o errexit -o pipefail -o nounset

cd $(dirname "$BASH_SOURCE")

rm -rf root

mkdir -p root/cnb/lifecycle
for file in ../out/lifecycle/*; do
  cp $file root/cnb/lifecycle/$(basename $file).exe
done

cat > root/cnb/order.toml <<EOF
[[order]]

  [[order.group]]
    id = "fake-buildpack-id"
    version = "0.0.0"
EOF

mkdir -p root/cnb/buildpacks/fake-buildpack-id/0.0.0
cat > root/cnb/buildpacks/fake-buildpack-id/0.0.0/buildpack.toml <<EOF
api = "0.2"

[buildpack]
id = "fake-buildpack-id"
name = "Fake"
version = "0.0.0"

[metadata]
include_files = ["bin/build.exe","bin/detect.exe","buildpack.toml"]

[[stacks]]
id = "io.buildpacks.stacks.windows1809"
EOF

cat > root/cnb/stack.toml <<EOF
[run-image]
image = "mcr.microsoft.com/windows/servercore:ltsc2019"
EOF

mkdir -p root/cnb/buildpacks/fake-buildpack-id/0.0.0/bin
GOOS=windows go build -o root/cnb/buildpacks/fake-buildpack-id/0.0.0/bin/detect.exe fake-buildpack/bin/detect/main.go
GOOS=windows go build -o root/cnb/buildpacks/fake-buildpack-id/0.0.0/bin/build.exe fake-buildpack/bin/build/main.go

mkdir -p root/workspace

source docker-env.sh

docker build -f Dockerfile --tag lifecycle-windows root