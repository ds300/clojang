## Clojang

Clojang is an attempt to write a Clojure interpreter in Go. The word 'Clojang' is pronounced however you like and is a bad portmanteau of 'Clojure' and 'golang'.

There are a few benefits in doing this, beyond my personal entertainment:

- It should be able to achieve fast start-up times, which might make using clojure for shell scripting something of a joy at last.
- The benefits of core.async's `go` macro come for free in the shape of Go's first-class goroutines and channels.

There are also some drawbacks:

- It will likely be slow. In fact, I've made a conscious descision to avoid non-trivial optimization strategies for the first iteration of this project. The interpreter will be designed for simplicity and functionality; if it's faster than a Turing machine made of legos, great! If not, at least I'll have learned something. Start-up time will be the only area in which optimization is planned to occur.
- Interop can be from the Go side only, and Go doesn't have dynamic linking so if you write wrapper code for a go library, you need to compile Clojang with your wrapper included. i.e. While Clojang doesn't have to reinvent the world, it does have to wrap and link the world. There may be an interesting way to deal with the dynamic linking issue; which involves having multiple copies of Clojang around, each with a different set of libs compiled in. I'd guess that 99% of the nasty organisation stuff can be taken care of automatically, but it's only worth exploring if Clojang manages to win hearts and minds.


### What's done?

It is very early days. Due to my inexperience, I'm prodding my way into the project from all angles to get a feel for how things will fit together. The only somewhat-solid code is the persistent hash map/set implementations, for which I've followed Rich Hickey by using Phil Bagwell's Hash Array Mapped Tries.


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

### | (• ◡•)| (❍ᴥ❍ʋ)
