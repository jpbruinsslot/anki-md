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
    -i, -input [input-file]
    -o, -output [output-file]
    -h, -help
```

Credits
-------

Sources that helped me write this:

- https://blog.gopheracademy.com/advent-2014/parsers-lexers/
- http://blog.leahhanson.us/post/recursecenter2016/recipeparser.html
- https://github.com/influxdata/influxql
- https://github.com/benbjohnson/sql-parser
