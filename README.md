# 📚 anki-md

A markdown to anki flash cards converter

## 📦 Installation

### Binary installation

[Download](https://github.com/jpbruinsslot/anki-md/releases) a
compatible binary for your system. For convenience, place `anki-md` in a
directory where you can access it from the command line.

### Via Go

```bash
$ go install github.com/jpbruinsslot/anki-md
```

## 💻 Usage

### Command line usage

```
NAME:
    anki-md - markdown to anki flash cards converter

USAGE:
    anki-md -i [input-file] -o [output-file]

EXAMPLES:

    $ anki-md -i deck.md -o deck.csv

    $ cat deck.md | anki-md -o deck.csv

    $ anki-md -i deck.md > test.csv

VERSION:
    0.1.0

WEBSITE:
    https://github.com/jpbruinsslot/anki-md

GLOBAL OPTIONS:
    -i, -input [input-file]     input file
    -o, -output [output-file]   output file
    -html                       convert field content to html
    -h, -help
```

### Deck creation

Create your deck, cards and fields as follows:

```
%% Who wrote the book "The C Programming Language"?

%% Brian W. Kernighan and Dennis M. Ritchie

---

%% Create a hello world program

%%

``
#include <stdio.h>

int main()
{
    printf("hello, world\n");
}
``

```

| identifier | explanation                                             |
| ---------- | ------------------------------------------------------- |
| `%%`       | represents a field, you can use multiple field per card |
| `---`      | represent a card                                        |
