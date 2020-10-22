# funct

Super simple functional library without crazy reflection magic.

# why?

There are a few libraries floating around for [golang](https://golang.org), but most *(if not all)* use crazy
reflection to evaluate types. `funct` does not need that, it's all done by manually implementing the standard types
individually. While this method does have a lot of redundancy, and it will be cleaned up with generics, this is a 
"necessary" evil.

# example

Code examples are a good way to show-case usability, so maybe this will help show the brilliance of this library.

Below is an example of turning JSON into YAML (sort of):

```go
package main

import (
	"fmt"
	"github.com/eyebrow-fish/funct"
)

func main() {
	var json []byte // pretend this is from json.Unmarshal(...)
	yaml := funct.ByteArray(json).Filter(func(i byte) bool { return i != '{' && i != '}' }).Join()
	fmt.Println(yaml) // this is basically what yaml is
}
```

# todo

- Tests
- Tests
- More Tests

# disclaimer

Progress is slow, but sporadic because of my job.
