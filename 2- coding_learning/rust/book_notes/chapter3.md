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

Like with immutable variables, **constant** are values bounded to a name and it not changable. But there are differences between const and variables:

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

## Scalar Types

- Scalar types represent a single value in Rust.
- Rust has 4 primary scalar types: int , float , bool , and char.

### Integer Types

- Int is a number without a decimal point. It uses the u32 type by default but it can be changed to i- or u- types depending on the size of the number.

#### Integer Types in Rust

- Length Signed Unsigned
- 8-bit i8 u8
- 16-bit i16 u16
- 32-bit i32 u32
- 64-bit i64 u64
- 128-bit i128 u128
- arch isize usize

#### Number literals can be written as

- Decimal: 98_222
- Hex: 0xff
- Octal: 0o77
- Binary: 0b1111_0000
- Byte (u8 only): b'A'

#### Integer Overflow

- Depending by the length of the integer, the compiler can throw an error saying that the calculation result is out of range e.g. 256 is out of range for u8 (1-255).

- To fix this issue, I can use the following methods provided by the standard library for primitive types:

- 1. Wrap all modes with the wrapping_* methods, such as wrapping_add or wrapping_mul .
- 2. Return None with the checked_* methods, such as checked_add or checked_mul .
- 3. Return the value and a boolean indicating whether there was an overflow with the overflowing_* methods, such as overflowing_add or overflowing_mul methods.
- 4. Saturate at the value’s minimum or maximum values with the saturating_* methods.

### Floating-Point Types

- Float is a number with a decimal point. It uses the f64 type by default but it can be changed to f32 type.
- f64 is the default type because on modern CPUs it’s roughly the same speed as f32 but is capable of more precision.
- f32 has single precision and f64 has double precision.

``` rust
fn main() {
    let x = 2.0; // f64

    let y: f32 = 3.0; // f32
}
```

### Numeric Operations

- The let statement allows mathematical operations to be performed on the variables.

``` rust
fn main() {
    let sum = 5 + 10; // sum
    let difference = 95.5 - 4.3; // subtraction
    let product = 4 * 30; // multiplication
    let quotient = 56.7 / 32.2;   // division
    let remainder = 43 % 5;  // remainder
}
```

### The Boolean Type

- Boolean type is a type with two possible values: true and false.
- In rust the boolean type is written as bool .

``` rust

fn main() {
    let t = true;

    let f: bool = false; // with explicit type annotation
}
```

- The method is used to check if the value is true or false. If statements can be used to check if the value is true or false.

### Character Type

- The char type is used to represent a single character.

``` rust
fn main() {
    let c = 'z';
    let z = 'ℤ';
    let heart_eyed_cat = '😻';
}
```

## Compound Types

- Compound types can group multiple values into one type.

### Tuple Type

- A tuple is a general way of grouping together a number of values with a variety of types into one compound type.

- A turple is created by writing a comma-separated list of values inside parentheses. Each position in the turple is set like an array.

``` rust
fn main() {
    let tup: (i32, f64, u8) = (500, 6.4, 1);

    let (x, y, z) = tup;

    println!("The value of y is: {y}");

    println!("The value of z is: {z}");
}

// This program create a tuple holding 3 values. 
// Then it binds the values to the variables x, y, and z.
// Then it prints the value of y and z.
```

- Turples can also be accessed by using a period (.) followed by the index of the value.

``` rust
fn main() {
    let tup: (i32, f64, u8) = (500, 6.4, 1);

    let five_hundred = tup.0;

    let six_point_four = tup.1;

    let one = tup.2;

    println!("The value of five_hundred is: {five_hundred}");

    println!("The value of six_point_four is: {six_point_four}");

    println!("The value of one is: {one}");

// This program create a tuple holding 3 values, and sets the values to the variable 500, 6.4, and 1.
// Then it prints the values based on the order of the index.
```

### Array Type

- An array is a collection of elements of the same type. Unlike a tuple, every element of an array must have the same type.
- An array is created by writing a comma-separated list of values inside square brackets. Each position in the array is set like an array.
- An array is an universal data type that can be used in any programming language.

``` rust

fn main() {
    let a = [1, 2, 3, 4, 5];

    let months = ["January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"];
}
```

``` python
# Array in Python
months = ["January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"]
```

``` Java
// Array in Java
String[] months = {"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"};
```

- In Rust, arrays can be set to a specific length. The length of the array is fixed and cannot be changed.

``` rust
fn main() {
    let a: [i32;5] = [1, 2, 3, 4, 5];
  
 // The array is signed to be an integer 32 bit and the length is 5. 
}
```

- In Rust, arrays can also be initialized with the same value for each element.

``` rust
fn main() {
    let a = [3; 5];
}

// This is an array that contains 5 elements, and each element is set to 3. Thats the same as writing [3, 3, 3, 3, 3].
```

### Accessing Array Elements

- An array can be defined being a single chunk of memory. In Rust an array element can be accessed by using the index of the element.

``` rust
fn main() {
    let a = [1, 2, 3, 4, 5];

    let first = a[0];
    let second = a[1];
}

// In thios program, the array a is defined to have 5 elements. 
// - The first element is accessed by using the index 0.
// - The second element is accessed by using the index 1.
```

### Invalid Array Element Access

- If an invalid index is used to access an element of an array, the program will panic and exit.

``` rust

// Chapter 2 Programming a Guessing Game 

use std::io;

fn main() {
    let a = [1, 2, 3, 4, 5];

    println!("Please enter an array index.");

    let mut index = String::new();

    io::stdin()
        .read_line(&mut index)
        .expect("Failed to read line");

    let index: usize = index
        .trim()
        .parse()
        .expect("Index entered was not a number");

    let element = a[index];

    println!("The value of the element at index {index} is: {element}");
}
```

- The program will compile and run, but when the user enters an index that is outside the valid range, the program will panic and exit.

``` bash
thread 'main' panicked at 'index out of bounds: the len is 5 but the index is 10', src/main.rs:19:19
note: run with `RUST_BACKTRACE=1` environment variable to display a backtrace
```

- Bash will generate this runtime error when the user enters an index that is outside the valid range. The program will panic and exit.
