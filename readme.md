## Clojang

Clojang is an attempt to write a Clojure interpreter in Go. The word 'Clojang' is pronounced k-low-yang and is a bad portmanteau of 'Clojure' and 'golang' (the much-better 'gojure' is already taken, but the project seems inactive).

There are a few benefits in doing this, beyond my personal entertainment:

- It should be able to achieve extremely fast start-up times, which might make using clojure for shell scripting something of a joy at last.
- It can piggyback on the Go runtime's garbage collection and thread management which I'm led to believe are good.
- Asynchronous channels will be a first-class citizens; no need to port that (admittedly quite impressive) hunk of code that is the `go` macro in core.async.

There are also some drawbacks:

- It will likely be slow. Like, slower than ruby slow. This is mostly because I'm a total Go newbie --- indeed, this is the first time I've written anything which compiles straight to machine code --- and I've only written one utterly trivial interpreter before now. However, even if I were an expert, Go is certainly going to struggle vs C performance-wise. I'd be interested in hearing from people in-the-know about what it would take to get any kind of thing resembling speed up in here.
- Interop can be from the Go side only, and Go doesn't have dynamic linking so if you write wrapper code for a go library, you need to compile Clojang with your wrapper included. So while Clojang doesn't have to reinvent the world, it does have to wrap and link the world.



### What's done?

- Implementation of Bagwell's HAMT for maps and sets, immutable.
- The beginnings of a reader

Bear with me, I've only beed doing this for a few days.
