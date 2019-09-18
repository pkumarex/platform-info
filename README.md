# ISecL platform-info

This library is used to get the platform related information such as hardwareUUID, bios name, CPU flags etc.
cmd - contains the command line utility for the library.

## System Requirements
- RHEL 7.5/7.6
- Epel 7 Repo
- Proxy settings if applicable

## Software requirements
- git
- Go 11.4 or newer

# Step By Step Build Instructions

### Install `go 1.11.4` or newer
The `platform-info` requires Go version 11.4 that has support for `go modules`. The build was validated with version 11.4 version of `go`. It is recommended that you use a newer version of `go` - but please keep in mind that the product has been validated with 1.11.4 and newer versions of `go` may introduce compatibility issues. You can use the following to install `go`.
```shell
wget https://dl.google.com/go/go1.11.4.linux-amd64.tar.gz
tar -xzf go1.11.4.linux-amd64.tar.gz
sudo mv go /usr/local
export GOROOT=/usr/local/go
export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
```

## Build platform-info

- Git clone the platform-info
- Run scripts to build the platform-info

```shell
git clone https://github.com/intel-secl/platform-info.git
cd platform-info
go build ./...
```

# Links
https://01.org/intel-secl/
