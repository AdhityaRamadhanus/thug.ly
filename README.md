# thug.ly
[![Go Report Card](https://goreportcard.com/badge/github.com/AdhityaRamadhanus/thug.ly)](https://goreportcard.com/report/github.com/AdhityaRamadhanus/thug.ly)

Thug life 24/7

Inspired by [thuglify-node](https://github.com/reyhansofian/thuglify-node) with some heuristic addition to calculate eyes region in case eyes not detected

<p>
  <a href="#installation">Installation |</a>
  <a href="#usage">Usage</a> |
  <a href="#licenses">License</a>
  <br><br>
  <blockquote>
	Thuglify your image
  </blockquote>
</p>

Installation
------------
* git clone
* go get -v
* make // it will create a bin folder and an executable called thugly

Usage
------------
```
thugly --input <image path> --output=<output path, default to thuglify.jpg>
```

```
NAME:
   thuglify - Thuglify your photos

USAGE:
   thuglify [global options] command [command options] [arguments...]

VERSION:
   1.0.0

AUTHOR:
   Adhitya Ramadhanus <adhitya.ramadhanus@gmail.com>

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --input value   Image input path
   --output value  Image output path (default: "thuglify.jpg")
   --help, -h      show help
   --version, -v   print the version

```

License
----

MIT © [Adhitya Ramadhanus]