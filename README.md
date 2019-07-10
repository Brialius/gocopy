# gocopy
## Usage
```
Usage of gocopy:
  -from string
        file to read from
  -limit int
        block size to copy (default -1)
  -offset int
        offset in input file (should be >= 0)
  -to string
        file to write to
  -v    verbosity mode
```
## Build
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
|release|set git tag and push to repo `make release ver=v1.2.3`|