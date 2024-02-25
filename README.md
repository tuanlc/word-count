# Book management

[![Go Report Card](https://goreportcard.com/badge/github.com/tuanlc/word-count)](https://goreportcard.com/report/github.com/tuanlc/word-count)
[![MIT license](https://img.shields.io/badge/license-MIT-brightgreen.svg)](https://opensource.org/licenses/MIT)
[![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-blue?logo=go&logoColor=white)](https://pkg.go.dev/github.com/tuanlc/word-count)

## Motivation

https://codingchallenges.fyi/challenges/challenge-wc

## Commands

### Run unit test

```bash
$ make test
```

### Run build

```bash
$ make build
```

### Clean folder

```bash
$ make clean
```

## Testing

`ccwc` is a tool that tries to clone features from the [Unix command](<https://en.wikipedia.org/wiki/Wc_(Unix)>) `wc`. Thus, the unit tests will use the output from `wc` as the source of truth. Moreover, the outcome of count words method could be different on different environments.

## License

[MIT](https://opensource.org/licenses/MIT)

Copyright (c) 2024-present, Tuan LE CONG
