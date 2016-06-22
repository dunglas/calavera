# Calavera, a static Single Page Application generator

Calavera helps to create a full-Javascript webapp from [Markdown](http://daringfireball.net/projects/markdown/) files.

Basically, it converts Markdown files stored in a [Git](https://git-scm.com/) repository to [JSON-LD](json-ld.org) files
using the [Schema.org](https://schema.org/) vocabulary.
Those files include metadata extracted from the Git repository (author name, last modification date...).

JSON-LD files are easily consumed by JavaScript libraries and framework such as [Angular](https://angular.io/), [Flux](https://facebook.github.io/flux/)
and [jQuery](https://jquery.com/).

Websites generated with Calavera can easily (and freely) be hosted on [GitHub Pages](https://pages.github.com/).

[![Build Status](https://travis-ci.org/dunglas/calavera.svg?branch=master)](https://travis-ci.org/dunglas/calavera)

## Install

The easiest way to install Calavera is to use the official [Docker](https://www.docker.com) image:

    $ docker pull dunglas/calavera

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

    $ docker run dunglas/calavera input_directory output_directory

Markdown files from the `input_directory` will be converted to JSON-LD files in `output_directory`.

### Options

* **-prettify**: Prettify json-ld output files

## License

Calavera is distributed under [the MIT license](LICENSE).

## Credits

Written in Go (golang) by [Kévin Dunglas](https://dunglas.fr) and sponsored by [Les-Tilleuls.coop](https://les-tilleuls.coop).

