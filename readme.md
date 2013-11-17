## Clojang

Clojang is an attempt to write a Clojure interpreter in Go. The word 'Clojang' is pronounced k-low-yang and is a bad portmanteau of 'Clojure' and 'golang' (the much-better 'Gojure' is already taken).

There are a few benefits in doing this, beyond my personal entertainment:

- It should be able to achieve extremely fast start-up times, which might make using clojure for shell scripting something of a joy at last.
- It can piggyback on the Go runtime's garbage collection and thread management which I'm led to believe are good.
- Asynchronous channels will be first-class citizens; no need to port that (admittedly very impressive) hunk of code that is the `go` macro in core.async.

There are also some drawbacks:

- It will likely be slow. Like, slower than ruby slow. This is mostly because I'm basically an idiot compared to people who write good fast VM-type things, and I've never read any books or academic papers about optimizing compilers. I might change the latter situation when it comes time to start writing the actual interprety bits.
- Interop can be from the Go side only, and Go doesn't have dynamic linking so if you write wrapper code for a go library, you need to compile Clojang with your wrapper included. i.e. While Clojang doesn't have to reinvent the world, it does have to wrap and link the world. There may be an interesting way to deal with the dynamic linking issue; which involves having multiple copies of Clojang around, each with a different set of libs compiled in. I'd guess that 99% of the nasty organisation stuff can be taken care of automatically. But I suppose that's only worth exploring if Clojang manages to win hearts and minds.


### What's done?

- 'Persistent rendition' of Bagwell's Hash-Array-Mapped-Tries for maps and sets.
- The beginnings of a reader

Bear with me, I've only been doing this for a few days.



### License

The MIT License (MIT)

Copyright © 2013 David Sheldrick

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.

## | (• ◡•)| (❍ᴥ❍ʋ)
