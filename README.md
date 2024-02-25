# Book management

[![MIT license](https://img.shields.io/badge/license-MIT-brightgreen.svg)](https://opensource.org/licenses/MIT)
[![Project CI](https://github.com/tuanlc/word-count/actions/workflows/ci.yml/badge.svg)](https://github.com/tuanlc/word-count/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/tuanlc/word-count)](https://goreportcard.com/report/github.com/tuanlc/word-count)
[![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-blue?logo=go&logoColor=white)](https://pkg.go.dev/github.com/tuanlc/word-count)

## Motivation

https://codingchallenges.fyi/challenges/challenge-wc

## Commands

### Development

- Run unit test

```bash
$ make test
```

- Run build

```bash
$ make build
```

- Clean folder

```bash
$ make clean
```

### Usage

#### Directly usage

After building the tool successfully, there will be the executable file `ccwc` in the root folder.

- Count bytes

```bash
$ ./ccwc -c test.txt
342190 test.txt
```

- Count lines

```bash
$ ./ccwc -l test.txt
7145 test.txt
```

- Count words

```bash
$ ./ccwc -w test.txt
58164 test.txt
```

- Count characters

```bash
$ ./ccwc -m test.txt
339292 test.txt
```

- Count characters

```bash
$ ./ccwc test.txt
7145 58164 342190 test.txt
```

#### Read from standar input

- Count bytes

```bash
$ cat test.txt | ./ccwc -c
342190 test.txt
```

- Count lines

```bash
$ cat test.txt | ./ccwc -l
7145 test.txt
```

- Count words

```bash
$ cat test.txt | ./ccwc -w
58164 test.txt
```

- Count characters

```bash
$ cat test.txt | ./ccwc -m
339292 test.txt
```

- Count characters

```bash
$ cat test.txt | ./ccwc
7145 58164 342190 test.txt
```

## Testing

`ccwc` is a tool that tries to clone features from the [Unix command](<https://en.wikipedia.org/wiki/Wc_(Unix)>) `wc`. Thus, the unit tests will use the output from `wc` as the source of truth. Moreover, the outcome of count words method could be different on different environments.

## License

[MIT](https://opensource.org/licenses/MIT)

Copyright (c) 2024-present, Tuan LE CONG
