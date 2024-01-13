# Strings

- Strings in go are all unicode
    - unicode is a particular technique thta allows us to represent international characters
    - ASCII represents only the characters from american english, and all the characters fit into one byte
    - Now when we moved to international languages(likes of arabic, chinese etc.) that needed different representation. UNICODE is a way to represent them.

- byte is a synonym for uint8
- Rune is a go equivalent of what you think of as a character (int32 for characters), 32bits because, 4 byte is enough to represent any UNICODE character.
- In order to make programs efficient, we don't want to represent every character everytime with 4 bytes, so there is a technique for encoding unicode called UTF-8, It is a short way of representing unicode in bytes, which actually co-incidentlly invented by guys who also worked on GO years ago at Bell Labs, as an efficient way to encode UNICODE

> So when we think of strings in go, physically they are the UTF-8 encoding of unicode characters.

- There's another type in go, called `byte` : It is just a synonym for an 8-bit integer.

> A string is physically a sequence of bytes that are required to encode the unicode chracters that are logically the Runes, and the rune is ofcourse just a synonym for a 32-bit integer.

- Length of string is the length of the byte string that's necessary to encode the string in UTF-8,
visually `élite` has only 5 characters but the length comes out to be 6, cause physically it is 6 bytes in UTF-8 encoding
- `[]byte` representation is [195 169 108 105 116 101], which has 6 values.

- So, length of string is the number of bytes required to represent the unicode characters not the number of unicode characters.

## String Structure

**The internal string representation is a pointer and a length**

s := "Hello, World"

hello := s[:5]
World := s[7:]

`s` is actually a string descriptor, which stores a pointer which points to the memory address of where in memory the bytes are stored and its byte length.
```
s
----------------------
|   ptr: addr        |
|   len: 12          |
----------------------
```

here, `hello` and `World` are substrings, and points to the memory address that is part of the string `s`. 
```
-----H e l l o ,   W o r l d-----
     ^
     |
    Pointer in 'hello' string, pointing to this address in 's'. It will also have its byte length stored in it.
``````

When we do, `s += 'es'` , we create a new copy, with different memory address.


Strings are passed by reference, thus they aren't coppied.


## string functions

Package `strings` has many functions on strings.

s := "a string"

x := len(s)

strings.Contains(s, "g")            // returns true
strings.Contains(s, "x")            // returns false

strings.HasPrefix(s, "a")           // returns true
strings.Index(s, "string")          // returns 2

s = strings.ToUpper(s)              // returns "A STRING"

> Note: we assign the result of ToUpper back to `s`, so, it's a copy, means a new string is created, value stored and now `s` is passed the memory address of this new string.
Go is a garbage collected language, when previous `s` is not reclaimed, it gets collected and discarded.

- We can't modify strings inplace, since they are immutable.


In `bufio`: Both `NewReader` and `NewScanner` use buffering to optimize I/O operations, but `NewScanner` adds the additional functionality of automatic tokenization based on delimiters, making it convenient for certain use cases like line-by-line reading.

> Note: In computing, a buffer is a temporary storage area that holds data while it's being transferred between two locations or processed by a program.
The default buffer size is typically 4 KB, but it can be adjusted.

Advantages of using Buffer:

- Using a buffer allows you to read a chunk of data at once, reducing the number of system calls and potentially improving performance

- Reading data in larger chunks can be more efficient when dealing with slow I/O operations, such as reading from a network or disk.

- Buffers help amortize the cost of these slow operations by allowing your program to read larger chunks at once and then process the data in memory.

- Allows for optimizations like lookahead and adaptive buffering for specific use cases.

    <u>Lookahead and Adaptive Buffering:</u>

    - Lookahead:

        - Pre-fetches data into the buffer, anticipating future read requests.
        - Reduces system calls and improves performance, especially for sequential reads.

    - Adaptive Buffering:

        - Dynamically adjusts buffer size based on reading patterns.
        - Balances memory usage for different use cases.


**NewReader:**

Purpose: Creates a buffered reader that efficiently reads data from an underlying io.Reader in chunks.

- Buffers data to reduce the number of system calls.
- Doesn't automatically handle tokenization or parsing.
- Provides methods like `Read`, `ReadByte`, `ReadLine`, etc., for basic reading operations.

When to Use:

- Reading raw data from sources like files, network connections, or strings.
- Basic reading tasks without needing tokenization.
- When you want more control over the reading process.


**NewScanner:**

Purpose: Creates a scanner that reads data from an underlying io.Reader and breaks it into tokens based on a delimiter.

Key Features:

- Automatically splits data into tokens (usually lines).
- Provides methods like `Scan`, `Text`, `Bytes`, etc., for token-based reading.

- Can customize the delimiter for tokenization.


When to Use:

- Reading line-by-line from text files or other structured data.

- Processing data in a token-based manner (e.g., parsing CSV, logs, configuration files).

- Handling newline-delimited data more conveniently.
