# Rust Book Notes
-----------------------------------------------------------
### Chapter 1

Rust is a low-level language that uses concepts to be easy to follow and understand.

Functions are coded like this:

``` rust
fn main() {
    println!("Hello, world!");
}
```
## How to Run normal Rust code: 

Run on the terminal with this code:

``` bash
rustc main.exe
./main
```
# Cargo

Cargo is the package manage of Rust. Most people install Cargo to build basic Rust code because it helps build the code better.


1. We can create a project using ***cargo new***.
2. We can build a project using ***cargo build***.
3. We can build and run a project in one step using cargo run.
4. We can build a project without producing a binary to check for errors using cargo check.
5. Instead of saving the result of the build in the same directory as our code, Cargo stores it in the target/debug directory.

-----------------------------------------------------------
### Chapter 2 
* Building a Guessing Game

``` rust 
use std::io; // import the std from the io library

fn main() {
    println!("Guess the number!");

    println!("Please input your guess.");

    let mut guess = String::new();

    io::stdin()
        .read_line(&mut guess)
        .expect("Failed to read line");

    println!("You guessed: {guess}");
}

// it prints two (2) outputs. Asking for the user input
// It stores the variable in a changable (mut) instance for later usage.
// Prints the result backs to the user.
``` 

















