# Proto

This is a repository for the protobuf files in the kohrVid auth project.

<!-- vim-markdown-toc GFM -->

* [Setup](#setup)
* [Mockgen](#mockgen)

<!-- vim-markdown-toc -->

## Setup

Protobuf files can be found in `./pb`. If changes are made to them, the `protoc`
command will need to be re-run.

Firstly, ensure that you have protoc installed. The best way to install this on
my machine was with the following:

    PROTOC_ZIP=protoc-3.7.1-linux-x86_64.zip
    curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v3.7.1/$PROTOC_ZIP
    sudo unzip -o $PROTOC_ZIP -d /usr/local bin/protoc
    sudo unzip -o $PROTOC_ZIP -d /usr/local 'include/*'
    rm -f $PROTOC_ZIP

To regenerate the go code associated with the protobuf files, run:

    protoc -I . --go_out=plugins=grpc:. ./*.proto

## Mockgen

The Mockgen tool in [GoMock](https://github.com/golang/mock#running-mockgen) has
been used to generate a mock client for use in client-side tests. The code for
this can be found in `./mock_proto`

If any changes are made to protobuf, most probably this will also need to be
regenerated. The command used to generate the existing code was as follows:

    mockgen github.com/kohrVid/auth/proto AuthenticationServiceClient > mock_proto/protoMock.go
