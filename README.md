anki-md 
=======

A markdown to anki flash cards converter

Installation
------------

#### Via Go

```bash
$ go get -u github.com/erroneousboat/anki-md
```

Usage
-----

#### Command line usage:

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
    https://github.com/erroneousboat/anki-md

GLOBAL OPTIONS:
    -i, -input [input-file]     input file
    -o, -output [output-file]   output file
    -html                       convert field content to html
    -h, -help
```

#### Deck creation

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
|------------|---------------------------------------------------------|
| `%%`       | represents a field, you can use multiple field per card |
| `---`      | represent a card                                        |

Credits
-------

Sources that helped me write this:

- https://blog.gopheracademy.com/advent-2014/parsers-lexers/
- http://blog.leahhanson.us/post/recursecenter2016/recipeparser.html
- https://github.com/influxdata/influxql
- https://github.com/benbjohnson/sql-parser
