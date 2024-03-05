default: test

# -timeout 	timeout in seconds
#  -v		verbose output
test:
	@ echo "+ $@"
	@ go test -timeout=5s -v

# `CGO_ENABLED=0`
# Because of dynamically linked libraries, this will statically compile the
# app with all libraries built in. You won't be able to cross-compile if CGO
# is enabled. This is because Go binary is looking for libraries on the
# operating system it’s running in. We compiled our app, but it still is
# dynamically linked to the libraries it needs to run
# (i.e., all the C libraries it binds to). When using a minimal docker image
# the operating system doesn't have these libraries.
#
# `GOOS=linux`
# We're setting the OS to linux (in case someone builds the binary on Mac or
# Windows)
#
# `-a`
# Force rebuilding of package, all import will be rebuilt with cgo disabled,
# which means all the imports will be rebuilt with cgo disabled.
#
# `-installsuffix cgo`
# A suffix to use in the name of the package installation directory
#
# `-o`
# Output
#
# `./bin/anki-md`
# Placement of the binary
#
# `.`
# Location of the source files
build:
	@ echo "+ $@"
	@ CGO_ENABLED=0 go build -a -installsuffix cgo -o ./bin/anki-md .

build-linux:
	@ echo "+ $@"
	@ GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -installsuffix cgo -o ./bin/anki-md-linux-amd64 .

build-darwin:
	@ echo "+ $@"
	@ GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -a -installsuffix cgo -o ./bin/anki-md-darwin-amd64 .


.PHONY: default test build build-linux build-darwin
