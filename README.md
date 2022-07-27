# go-scanner

This is a local network scanning tool to find hosts connected to your network

## why go

There are so many tools out to achieve same goal. we used golang to make scanning faster

## Install 

### Install with docker

1. install docker in your machine.
2. run `sudo docker pull nibrasmuhamed/go-scanner:1.1`
3. `docker run nibrasmuhamed/go-scanner:1.1 <args here>` // you may need to run as super user

### Build docker image with source code.

1. clone this repository
2. run `docker build -t <tag:version>`
3. `docker run <tag:version>`
    replace tag and version accordingly. 

### Use with source code.

1. clone this repository
2. install golang on your machine
3. executed `go run *.go`

### Use with build releases

1. download pre-compiled binary in releases tab.
2. execute. `./goscanner`
sudo docker pull nibrasmuhamed/go-scanner:1.1

## Usage

go-scanner built with spf13/cobra library.

`./go-scanner` 

### Sub commands

1. `./go-scanner completion` 
    Generate the autocompletion script for go-scanner for the specified shell.
    See each sub-command's help for details on how to use the generated script.

2. `./go-scanner scan`
    scans provided router to find other connected hosts.
    routerip/mask

    eg : 172.17.0.1/16

3. `./go-scanner help`
    prints manual for user.


