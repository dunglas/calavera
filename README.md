# Calavera, a static Single Page Application generator

Calavera helps to create a full-Javascript webapp from [Markdown](http://daringfireball.net/projects/markdown/) files.

Basically, it converts Markdown files stored in a [Git](https://git-scm.com/) repository to [JSON-LD](json-ld.org) files
using the [Schema.org](https://schema.org/) vocabulary.
Those files include metadata extracted from the Git repository (author name, last modification date...).

JSON-LD files are easily consumed by JavaScript libraries and framework such as [React](https://facebook.github.io/react/),
[Angular](https://angular.io/) or [jQuery](https://jquery.com/).

Websites generated with Calavera can be freely and easily hosted on [GitHub Pages](https://pages.github.com/).

[![Travis](https://img.shields.io/travis/dunglas/calavera.svg?maxAge=2592000)](https://travis-ci.org/dunglas/calavera)
[![Coveralls](https://img.shields.io/coveralls/dunglas/calavera.svg?maxAge=2592000)](https://coveralls.io/github/dunglas/calavera)
[![Docker Automated buil](https://img.shields.io/docker/automated/dunglas/calavera.svg?maxAge=2592000)](https://hub.docker.com/r/dunglas/calavera/)

## Install

The easiest way to use Calavera is through [Docker](https://www.docker.com). Just go to the usage section below.

Alternatively, you can compile Calavera from sources.

The [Go](https://golang.org/) programming language is the only required dependency to compile Calavera.

If you don't already have a [Go workspace](https://golang.org/doc/code.html#Workspaces), create it: 

    $ mkdir -p ~/workspace/go
    $ export GOPATH=~/workspace/go

Create the appropriate directory structure and download the source code:

    $ mkdir -p ~/workspace/go/src/github.com/dunglas
    $ cd ~/workspace/go/src/github.com/dunglas
    $ git clone git@github.com:dunglas/calavera.git

Go to the source code directory, download external libraries and compile the program:    

    $ cd calavera
    $ go get
    $ go install

## Usage

Using Docker:

    $ docker run -v /my/src/directory:/in -v /my/output/directory:/out dunglas/calavera /in /out

If you installed it from source:

    $ ~/workspace/go/bin/calavera input_directory output_directory

Markdown files from the `input_directory` will be converted to JSON-LD files in `output_directory`.

### Options

* **-prettify**: Prettify generated JSON-LD files

## License

Calavera is distributed under [the MIT license](LICENSE).

## Credits

Written in Go (golang) by [Kévin Dunglas](https://dunglas.fr) and sponsored by [Les-Tilleuls.coop](https://les-tilleuls.coop).
