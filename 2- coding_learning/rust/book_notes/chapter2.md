
### Chapter 2 
* Building a Guessing Game

#### Version 1.0 

``` rust 
use std::io; // import the std from the io library

fn main() {
    println!("Guess the number!");

    println!("Please input your guess.");

    let mut guess = String::new();

    io::stdin()
        .read_line(&mut guess)
        .expect("Failed to read line"); // what if the program fails/

    println!("You guessed: {guess}");
}

// it prints two (2) outputs. Asking for the user input
// It stores the variable in a changable (mut) instance for later usage.
// Prints the result backs to the user.
``` 
#### Version 1.1

``` rust 
use std::io;
use rand::Rng;

fn main() {
    println!("Guess the number!");

    let secret_number = rand::thread_rng().gen_range(1..=100);

    println!("The secret number is: {secret_number}");

    println!("Please input your guess.");

    let mut guess = String::new();

    io::stdin()
        .read_line(&mut guess)
        .expect("Failed to read line");

    println!("You guessed: {guess}");
}

// imports both the io library and the RNG library

// sets the sewcret number betweeen 1 and 100

// it prints the secret number. and ask for the usre input

// it creates a changable variable that reads and prints the user output.
```
### Version 1.2

``` rust 
use rand::Rng;
use std::cmp::Ordering;
use std::io;

fn main() {
    println!("Guess the number!");

    let secret_number = rand::thread_rng().gen_range(1..=100);

    println!("The secret number is: {secret_number}");

    println!("Please input your guess.");

    // --snip--

    let mut guess = String::new();

    io::stdin()
        .read_line(&mut guess)
        .expect("Failed to read line");

    let guess: u32 = guess.trim().parse().expect("Please type a number!");

    println!("You guessed: {guess}");

    match guess.cmp(&secret_number) {
        Ordering::Less => println!("Too small!"),
        Ordering::Greater => println!("Too big!"),
        Ordering::Equal => println!("You win!"),
    }
}
// import libraries
// creates a mutable varaible that prints if the result is empty based on another variable
// if the variable is too low prints "Too small"
// if is too big print "Too big"
// If the result matches the guess print ""You win""
```
It should print this 

``` bash
$ cargo run
   Compiling guessing_game v0.1.0 (file:///projects/guessing_game)
    Finished dev [unoptimized + debuginfo] target(s) in 0.43s
     Running `target/debug/guessing_game`
Guess the number!
The secret number is: 58
Please input your guess.
  76
You guessed: 76
Too big!
```
### Version 1.3

``` rust 
// prints library 
use rand::Rng;
use std::cmp::Ordering;
use std::io;

// main function for this program 

fn main() {  
    println!("Guess the number!"); // program title 
 
    let secret_number = rand::thread_rng().gen_range(1..=100); // sets the program guesses between 1 and 100

    loop {
        println!("Please input your guess."); // ask for the user input

        let mut guess = String::new();

        io::stdin()
            .read_line(&mut guess)
            .expect("Failed to read line"); // creates a variable if the user ask for the wrong number. And makes conditions if there's an error 

        let guess: u32 = match guess.trim().parse() {
            Ok(num) => num,
            Err(_) => continue,
        };

        println!("You guessed: {guess}"); // pritns the guesses based if the matches is far, close, or exactly as the program guess

        match guess.cmp(&secret_number) {
            Ordering::Less => println!("Too small!"),
            Ordering::Greater => println!("Too big!"),
            Ordering::Equal => {
                println!("You win!");
                break;
            }
        }
    }
}
```