#### Compiler & Language definition for the Ethereum project

This is the Mutan 2.0 branch.

Mutan is a C-Like language for the Ethereum project. Mutan supports a
full, dynamic higher level language that compiles to native
Ethereum Assembler. The language definition and documentation
can be found on [go-ethereum wiki](https://github.com/ethereum/go-ethereum/wiki/Mutan).

A simple online editor and compiler can be found [here](http://mut.etherian.io)

### Install compiler

`go get -u github.com/obscuren/mutan/mutan`

### CLI

```
mutan [flags] filename
-s=""        Compile string
-asm         Output asm instead of bytecode/hex
-b           Output raw byte array instead of hex
```

### Release notes

##### 0.5 (Upcoming)

* Scoping using `{ var x = 10; }`
* Proper else if clauses `if { } else if { } else { }`
* Added logical and / or `var x = 1 && 2; if 1 && 0 { }`
* Changed the `this` keyword to `message`.

##### 0.4

Didn't keep release notes :-)

### Syntax

```go
import "std.mu"

func fn(var a, var b) {
	var[2] c
	c[0] = a
	c[1] = b
	return c[1]
}

var a = fn(0, 1)
b := 10

if a > b {
    stop()
} else if 0 && 1 || 0 {
    // :-(
    stop()
} else {
    // :-)
    if !a {
        if message.data[0] ** 10 >= 10 {
            message.data[0] = 1000;
        }
    }
}

message.store[a] = 10000
message.store[b] = message.origin()

for i := 0; i < 10; i++ {
    var[10] out
    call(0xaabbccddeeff112233445566, 0, 10000, i, out)
}

// tx without input data
transact(0xa78f6abe, 10000, nil)
// no args and return values
call(0xab, 0, 10000, nil, nil)
// create contract
var ret = create(value, 0xaabbccddeeff0099887766552211)

asm {
    push1 10
    push1 0
    mstore
}

return 20
```

Mutan &copy; Jeffrey Wilcke
