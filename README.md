# Unix users and groups query REST service - nixug
###


Passwd as a Service
The idea of this challenge is to create a minimal HTTP service 
that exposes the user and group information on 
a UNIX-like system that is usually locked away in the UNIX /etc/passwd and /etc/groups files.
While this service is obviously a toy (and potentially a security nightmare), 
please treat it as you would a real web service. 
That means write production quality code per your standards, 
including at least: Unit Tests, and README documentation. 

To aid testing and deployment, the paths to the passwd and groups file 
should be configurable, defaulting to the standard system path. 
If the input files are absent or malformed, your service must indicate 
an error in a manner you feel is appropriate.
This service is read-only but responses should reflect changes made to 
the underlying passwd and groups files while the service is running. 


## Source code

Source code is available on github - https://github.com/ivost/nixug

git clone https://github.com/ivost/nixug.git

# Implementation

Linux users are stored in /etc/passwd file (but there are no passwords in it)
and groups - in /etc/groups.

For peculiar reasons user and group names and numerical ids may contain duplicate values.

I started with simple storage model using go map - but then decided to use array and 
allow for cannonical duplication.
In case of duplication linux uses the first match - so I am doing something similar.

The only perf improvement to help with O(n) when using arrays is to do some sorting of user names and user groups when storing them in the array.

Once sorted - search by name uses binary search which is O(log n).
Assumption is that querying by name is the most frequent operation.

### Caching and monitoring for file changes

nixug is using fsnotify to react on modifications of passwd and group files
In-memory cache is not updated on every change - but on the first releveant client request when "dirty" flag indicates changes.
 

## Prerequisites

make (required for Makefile convenience commands)

one of: 
* go v.1.11 or later
* docker

nixug uses go modules - it requires go v.1.11 or later and export GO111MODULE=on
alternatively it can be run via docker command - the image is available in dockerhub

It uses labstack echo web framework  https://godoc.org/github.com/labstack/echo


### Configuration

via config.json (expected in current directory when running)

example 

```bash
{
    "Host": "0.0.0.0",
    "Port": 8080,
    "UserFile": "/etc/passwd",
    "GroupFile": "/etc/group",
    "Auth": false
}

```
where:

Host and Port can be used to configure listening address and port

UserFile and GroupFile point to passwd and group files

Auth enables or disables API token authentication

### Running/building

Assuming make and docker are available the service can be started just by

```
# docker pull and run
make drun

# or without make (prefix with sudo if docker requires root access)
docker run --rm -d -n nixug -p 8080:8080 ivostoy/nixug:1.0.5.24

# demo users
make users

# demo groups
make groups

```

Running 
```
make help
```
will show available targets

```
build    - build ./nixug binary
install  - install nixug binary (assuming GOPATH is in PATH)
run      - build using go and run app 
auth     - get auth token

test     - unit tests (no output)
testv    - unit tests with output
testr    - unit tests with race detection
testi    - integration tests (will start nixug in background)

check    - race test + go vet + go fmt
scheck   - static analysis
pedantic - check unparam errcheck
bench    - run benchmark tests
cpu      - cpu profiling

drun     - docker run will pull/run the image to dockerhub
docker   - will build docker image
push     - will push the built image to dockerhub
kill     - will kill running nixus and running container

for demo - assumes running nixug

health   - health check
get-u    - get users 
get-g    - get groups
```

### Testing

#### Unit tests

make test

#### Integration tests

make testi

### Profiling (optional)

https://github.com/google/pprof/blob/master/doc/README.md

go get github.com/pkg/profile
go get github.com/google/pprof

go tool pprof http://localhost:6060/debug/pprof/trace?seconds=5
go tool pprof http://localhost:6060/debug/pprof/heap

nixug -cpu prof-file

pprof -http localhost:9090 prof-file

----

### Authentication (optional)

Uses JWT auth.token in the headers - can be turned on/off in configuration

 "Auth": true

#### How to demo/test

```
# installing httpie (optional)
pip install -U httpie-jwt-auth

http localhost:8080/auth/nix/nix

TOKEN=$(http localhost:8484/v1/auth/nix/nix); echo $TOKEN

http \
  --auth-type=jwt --auth=$TOKEN \
  localhost:8080/users/0

```
