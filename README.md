# bazelsample

Sample App for Bazel and Go Modules.

## Install Bazel

- https://docs.bazel.build/versions/master/install-os-x.html
  - binary install

## Setup Workspace

- https://github.com/bazelbuild/rules_go#setup
  - Using gazelle for generating build files

```
# Create WORKSPACE file
vim WORKSPACE

# Create BUILD.bazel file
vim BUILD.bazel

# Generate build files
bazel run //:gazelle
```

## Setup go modules

```
export GO111MODULE=on
go mod init

mkdir -p cmd/bazelsample
vim cmd/bazelsample/bazelsample.go
vim sample.go
vim sample_test.go

go test
go build -o bazelsample ./cmd/bazelsample
rm bazelsample
```

## Update dependencies from go.mod file

```
bazel run //:gazelle -- update-repos -from_file=go.mod
bazel run //:gazelle
```

## Test

```
bazel test //...
```

## Build

```
bazel build //cmd/bazelsample
```

# References

- http://mjhd.hatenablog.com/entry/2018/07/22/224528
- https://docs.bazel.build/versions/master/bazel-overview.html
- https://github.com/bazelbuild/rules_go
- https://github.com/bazelbuild/bazel-gazelle
