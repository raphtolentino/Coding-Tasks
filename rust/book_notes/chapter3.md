# Chapter 3

> Chapter 3.1 : Variable and Multability

Variables in Rust is not **_multable_** meaning that once its variable its declared to memory it can't be changed.

For example the project file for this project is:

``` rust
fn main() {
  let x = 5;
  println!("The value of x is: {x}");
  x = 6;
  println!("The value of x is: {x}");

  // this code will create a error because its trying to print the same undeclared variable two times. If changed to mut x it could be run.
}
```

``` bash
$ cargo run
   Compiling variables v0.1.0 (file:///projects/variables)
error[E0384]: cannot assign twice to immutable variable `x`
 --> src/main.rs:4:5
  |
2 |     let x = 5;
  |         -
  |         |
  |         first assignment to `x`
  |         help: consider making this binding mutable: `mut x`
3 |     println!("The value of x is: {x}");
4 |     x = 6;
  |     ^^^^^ cannot assign twice to immutable variable

For more information about this error, try `rustc --explain E0384`.
error: could not compile `variables` due to previous error
```

**The error "cannot assign twice to immutable variable" means that the program tried to assign two variable to a static variable.**

Changing the code to will allow the variable to be mutable:

```rust
let mut ...
```

The program will print this outcome:

``` rust
$ cargo run
   Compiling variables v0.1.0 (file:///projects/variables)
    Finished dev [unoptimized + debuginfo] target(s) in 0.30s
     Running `target/debug/variables`
The value of x is: 5
The value of x is: 6
```

> Chapter 3.1: Constants in Rust

*Like with immutable variables, ***constant*** are values bounded to a name and it not changable. But there are differences between const and variables:*

- First constansts are always immutable (mut are not allowed).
- Second, constants can be declared in any scope, including the global scope.
- Constant may be set only to a constant expression, not the result of a function call or any other value that could only be computed at runtime.

``` rust
const THREE_HOURS: u32 = 3 * 60 * 60;

// this code is for finding the amount of seconds in 3 hours.

// 60 seconds in a minute. 60 minutes in an hour. 3 hours.

```
> Chapter 3.1: Shadowing

- Shadowing is a way to declare a new variable with the same name as a previous variable.

- The first variable is shadowed by the second, which means that the first variable is effectively hidden.

- Shadowing is not multable, because we’re effectively creating a new variable when we use the let keyword again.

``` rust
fn main() {
    let x = 5;

    let x = x + 1;

    {
        let x = x * 2;
        println!("The value of x in the inner scope is: {x}");
    }

    println!("The value of x is: {x}");
}

// this code binds the variable x to a variable of 5 .
// then it binds the variable x to a variable of 6 .
// then it binds the variable x to a variable of 12 .
```

> Chapter 3.2 : Data Types

- Evevy value is part of an specific data type, which tells Rust how to work with that specific data.

- Rust is a static language, which means that the compiler must know all data types before compiling.