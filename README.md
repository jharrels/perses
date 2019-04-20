Perses is a CloudFront invalidator.  It is named after [the Greek titan god of the Destruction](https://en.wikipedia.org/wiki/Perses_(Titan)).
## Installation

First install [Go](http://golang.org).

If you just want to install the binary to your current directory and don't care about the source code, run

```bash
GOBIN="$(pwd)" GOPATH="$(mktemp -d)" go get github.com/jharrels/perses
```

## Usage

```bash
$ perses -h
Usage of Perses:

    perses [options] <invalidation path>...

Invalidation path defaults to '/*'.

AWS credentials taken from ~/.aws/ or from "AWS_ACCESS_KEY_ID", "AWS_SECRET_ACCESS_KEY", and other AWS 
configuration environment variables.


Options:

  -dist string
        CloudFront distribution ID
  -query string
        ClooudFront invalidation ID
  -ref string
        CloudFront 'CallerReference', a unique identifier for this invalidation request. (default: Unix timestamp)


$ perses -dist EABC123EFG4567
2019/04/19 14:35:43 Invalidation ID: "IQ2JXQ53AYXGBB"
```
