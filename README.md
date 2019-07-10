# gocopy
### make goals
|Goal|Description|
|----|-----------|
|setup|download and install required dependencies|
|test|run tests|
|build|build binary: `bin/gocopy` of `bin/gocopy.exe` for windows|
|install|install binary to `$GOPATH/bin`|
|lint|run linters|
|clean|run `go clean`|
|mod-refresh|run `go mod tidy` and `go mod vendor`|
|ci|run all steps needed for CI|
|version|show current git tag if any matched to `v*` exists|