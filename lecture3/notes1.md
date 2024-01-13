# Data Types

- For `int` the default is `int64`
- For floating point type, we just have `float32` and `float64`(default)

- Don't use floating points for monetary calculations, use GoMoney package.

- The 8 in `%8T` is specifying the minimum width of the printed output. It means that the type name will be padded with spaces on the left if it's less than 8 characters wide.

- The `%v` is a general-purpose formatting verb used to print the value.


## Special Types:

- `bool` has two values `true`, `false`
    - These are *not* convertible to/from integers!, means no way to convert a bool to an int and vice-versa


- `error`: a special type with one function, `Error()`
    - An error may be nil or non-nil

- Pointers are physically addresses, logically opaque
    - a pointer may be nil or non-nil
    - No pointer manipilation except through package unsafe


## Initializations

Go initializes all the variables to "zero" by default if you don't:
- There are no uninitialized variables in Go
- All numeriacal types get 0 ( float 0.0, complex 0i )
- bool gets false
- string gets "" (the empty string, length 0)

- Everything else gets nil:
    - pointers
    - slices
    - maps
    - channels
    - functions(function variables)
    - interfaces

- For aggregate types, all members get their "Zero" values


## Constants

- Only numbers, strings and booleans can be constants(immutable)

- Constant can be a literal or compile-time function of a contant
```
const(
    a= 1                    // int
    b= 2 * 1024             // 2048
    c= b << 3               // 16384

    g uint8 = 0x07          // 7
    h uint8 = g & 0x03      // 3

    s= "a string"
    t= len(s)               // 8
    u= s[2:]                // SYNTAX ERROR
)
```

- can't do `++n` in go, only `n++` is valid here

- To read out put from nums.txt, use command: `go run main.go < nums.txt`

