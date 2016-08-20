# Getting The Code

### Go Environment
If not set already, please set your `GOPATH` and `GOBIN` environment variables. For example:
```bash
mkdir -p ~/go
export GOPATH=~/go
export GOBIN=$GOPATH/bin
```

### Get the code
Use [`go`](https://golang.org/cmd/go/) tool to "get" (i.e. fetch and build) `mainflux-lite` package:
```bash
go get github.com/mainflux/mainflux-lite
```

This will download the code to `$GOPATH/src/github.com/mainflux/mainflux-lite` directory,
and then compile it and install the binary in `$GOBIN` directory.

Now you can run the server:
```bash
$GOBIN/mainflux-lite
```

Please note that you can pass the path to the configuration file `config.yml` as an argument.
If no paramter is passed, default path is $GOPATH/src/github.com/mainflux/mainflux-lite`.

Note also that using `go get` is prefered than out-of-gopath code fetch by cloning the git repo like this:
```
git clone https://github.com/Mainflux/mainflux-lite && cd mainflux-lite
go get
go build
./mainflux-lite ./config/config.yml
```
