## nixug


config file

### Uses go modules - requires go v.1.11 or later

after git pull - try first

go mod tidy

If invoked with -mod=readonly, the go command is disallowed from the implicit
automatic updating of go.mod described above. Instead, it fails when any changes
to go.mod are needed. This setting is most useful to check that go.mod does
not need updates, such as in a continuous integration and testing system.
The "go get" command remains permitted to update go.mod even with -mod=readonly,
and the "go mod" commands do not take the -mod flag (or any other build flags).


uses labstack echo
https://godoc.org/github.com/labstack/echo

https://godoc.org/github.com/labstack/echo?importers

https://engineering.checkr.com/introducing-checkrs-integration-testing-workflow-and-openmock-572c64209891


PROFILING 

https://github.com/google/pprof/blob/master/doc/README.md

go get github.com/pkg/profile
go get github.com/google/pprof

https://golang.org/pkg/net/http/pprof/
https://blog.golang.org/2011/06/profiling-go-programs.html

go tool pprof http://localhost:6060/debug/pprof/trace?seconds=5

go tool pprof http://localhost:6060/debug/pprof/heap


nixug -cpu prof-file

pprof -http localhost:8080 prof-file

----

pip install -U httpie-jwt-auth

http localhost:8080/v1/auth/nix/nix

TOKEN=$(http localhost:8484/v1/auth/nix/nix); echo $TOKEN

http \
  --auth-type=jwt --auth=$TOKEN \
  localhost:8080/v1/groups/1
