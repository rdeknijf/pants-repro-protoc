# `export-codegen` problem reproduction repo

This repo is an example that reproduces the following:

    11:11:48.11 [ERROR] 1 Exception encountered:

    Engine traceback:
    in `export-codegen` goal

    ProcessExecutionFailure: Process 'Build Go gRPC protobuf plugin for `protoc`.' failed with exit code 1.
    stdout:

    stderr:
    go: google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2.0: loading deprecation for google.golang.org/grpc/cmd/protoc-gen-go-grpc: module lookup disabled by GOPROXY=off



    Use `--keep-sandboxes=on_failure` to preserve the process chroot for inspection.


So, `export-codegen` doesn't work. Also, the [docs](https://www.pantsbuild.org/stable/docs/go/integrations/protobuf) describe that Pants auto-regenerates any protobuf output. But doesn't. If you remove the `pkg/person.pb.go` file, and run `pants test ::` tests fail. If you run `protoc --go_out=. --go_opt=paths=source_relative person.proto` (in `pkg`), tests are good again.