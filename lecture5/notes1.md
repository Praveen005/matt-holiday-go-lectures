# Arrays, Slices and Maps

## Arrays:

- Arrays are typed by size, which is fixed at compile time.
- Arrays are passed by value, thus elements are copied.

## Slices:

Don't skip this:

[Cheatsheet](https://ueokande.github.io/go-slice-tricks/)

- The make built-in function allocates and initializes an object of type slice, map, or chan (only).

- Slices have variable length, backed by some array; they are copied when they outgrow the array.

- when we suppose say, `var a []int`, `a` is a descriptor here.
- A slice descripter has a pointer, length, and capacity