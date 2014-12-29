hd
==

A small hex dumper utility, because I never remember how to get `od` + `dd`
to do what I want...

```
$ hd --help
Usage:
  hd [options] [filename]

Options:
  -lim=0: Limit read bytes
  -skip=0: Skip bytes at start

$ hd -lim 8 -skip 8 -lim 64 main.go
00000000  6d 61 69 6e 0a 0a 69 6d  70 6f 72 74 20 28 0a 09  |main..import (..|
00000010  22 65 6e 63 6f 64 69 6e  67 2f 68 65 78 22 0a 09  |"encoding/hex"..|
00000020  22 66 6c 61 67 22 0a 09  22 66 6d 74 22 0a 09 22  |"flag".."fmt".."|
00000030  69 6f 22 0a 09 22 6c 6f  67 22 0a 09 22 6f 73 22  |io".."log".."os"|
```
