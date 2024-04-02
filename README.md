# linereturn

Linter requires a new line before return statements if the previous line in a block is a real statement.

# Example

Examples of incorrect code:

```go
func fooa() int {
    o := []int{0, 1}
    return o[0]
}

func foob(s string) *a {
    o := &a{
     a: s,
    }
	
    return o
}

func fooc() int {
    defer bar()
    return 0
}

func fooe(s string) interface{} {
    o := foob(
        s,
    )
	
    return o
}
```

Examples of correct code:

```go
func fooa() int {
    o := []int{0, 1}
	
    return o[0]
}

func foob(s string) *a {
    o := &a{
        a: s,
    }
    return o
}

func fooc() int {
    defer bar()
    
    return 0
}

func fooe(s string) interface{} {
    o := foob(
        s,
    )
    return o
}
```

# Args

* `-block-size n` size of the block (including return statement that is still "OK") so no return split required.