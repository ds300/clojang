## Clojang

Clojang is an attempt to write a Clojure interpreter in Go. The word 'Clojang' is pronounced however you like and is a bad portmanteau of 'Clojure' and 'golang'.

There are a few benefits in doing this, beyond my personal entertainment:

- It should be able to achieve fast start-up times, which might make using clojure for shell scripting something of a joy at last.
- The benefits of core.async's `go` macro come for free in the shape of Go's first-class goroutines and channels.
- With no ties to the JVM it can provide stuff like TCO, throwable anythings, protocols all-the-way-down, nice stacktraces, a good debugging experience, and so on.

There are also some drawbacks:

- It will likely be slow. In fact, I've made a conscious descision to avoid non-trivial optimization strategies for the first iteration of this project. The interpreter will be designed for simplicity and functionality; if it's faster than a Turing machine made of legos, great! If not, at least I'll have learned something. Start-up time will be the only area in which optimization is planned to occur.
- Interop can be from the Go side only, and Go doesn't have dynamic linking so if you write wrapper code for a go library, you need to compile Clojang with your wrapper included. i.e. While Clojang doesn't have to reinvent the world, it does have to wrap and link the world. There may be an interesting way to deal with the dynamic linking issue, which will warrant investigation if Clojang manages to win hearts and minds.

### What's done?

It is very early days. Due to my inexperience, I'm gently prodding my way into the project from all angles to get a feel for how things will fit together. The only somewhat-solid code is the persistent hash map/set implementations, for which I've followed Rich Hickey by using Phil Bagwell's Hash Array Mapped Tries.

### Can I Contribute?

Only with ideas at the moment. Once there is an alpha version up and running, I'll be excited to start accepting pull requests and doing the whole OSS deal.

Contact me: ```(apply str (reverse [".com" "gmail" "@" "djsheldrick"]))```



### License

Copyright © 2013 David Sheldrick

Distributed under the Eclipse Public License, the same as Clojure.

### | (• ◡•)| (❍ᴥ❍ʋ)
