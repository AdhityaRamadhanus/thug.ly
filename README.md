# thug.ly
[![Go Report Card](https://goreportcard.com/badge/github.com/AdhityaRamadhanus/thug.ly)](https://goreportcard.com/report/github.com/AdhityaRamadhanus/thug.ly)

Thug life 24/7

Inspired by [thuglify-node](https://github.com/reyhansofian/thuglify-node) with some heuristic addition to calculate eyes region in case eyes not detected

<p>
  <a href="#installation">Installation |</a>
  <a href="#usage">Usage</a> |
  <a href="#results">Results</a> |
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
thugly --input <image path> --output=<output path, default to thuglify.jpg> --text <text to add at the bottom of image, like "Deal with it">
```

```
NAME:
   thugly - Thuglify your photos

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
   --text value    Text to add at the bottom
   --output value  Image output path (default: "thuglify.jpg")
   --help, -h      show help
   --version, -v   print the version

```

Results
------------
![adit-thuglify](https://cloud.githubusercontent.com/assets/5761975/24078484/837626f8-0ca1-11e7-9d53-5a0a0fe7cfdb.jpg)
![lena-thuglify](https://cloud.githubusercontent.com/assets/5761975/24078486/837e0b70-0ca1-11e7-8f39-b46301ff264d.jpg)
![hitler-thuglify](https://cloud.githubusercontent.com/assets/5761975/24257399/5e4fa230-101d-11e7-8fd3-20f4dab509fd.jpg)
![lenin-thuglify](https://cloud.githubusercontent.com/assets/5761975/24257400/5e8c8ad8-101d-11e7-9d31-2db30e8dcd92.jpg)
![obama-thuglify](https://cloud.githubusercontent.com/assets/5761975/24257401/5e92d2c6-101d-11e7-9eea-ff681665258b.jpg)

License
----

MIT © [Adhitya Ramadhanus]