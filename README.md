# funct

Super simple functional library without crazy reflection magic.

# why?

There are a few libraries floating around for [golang](https://golang.org), but most *(if not all)* use crazy
reflection to evaluate types. `funct` does not need that, it's all done by manually implementing the standard types
individually. While this method does have a lot of redundancy, and it will be cleaned up with generics, this is a 
"necessary" evil.

# disclaimer

Progress is slow, but sporadic because of my job.
