# qcat

Concatenate FILE(s) to standard output, quoting each line

With no FILE, or when FILE is -, read standard input.

Usage:

```
qcat [-l='"'] [-r='"'] [-escape=true] [-trim=true] [FILE]...
```

## FAQ

> Why?

Very occasionally I have the need to quote all lines in a file. I'm interested in
learning Go and thought it'd be a fun tool to write.

> Couldn't you just use something like `awk '{ print "\""$0"\""}' FILE`

Sure, but this is more fun.

> Where does the name come frome?

`qcat`, short for "quote cat".
