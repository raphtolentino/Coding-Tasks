# Chapter 4 - Understanding Ownership

> Chapter 4.1 - What is Ownership?

Ownership refers to a rule set that dictates for Rust manages memory.

- 1. Some program languages has a garbage collection program that runs together with the compiler that regurarly looks for non-longer-used memory as that the program is running.

- 2. Some program languages the programmer has to explicitly allocate the variables to free memory.

- 3. Rust is a bit different, it uses a ownership system with a set of rules that checks alongside of the company. If any rules are violeted the program will not run.

## The Stack and the Heap

### Part 1: Definitions

> The Stack

- Many different programing languages, this concept of stack and heap is not as relevant, but in Rust it's affects how the language behaves.

- Both the stack and the heap are part of memory avaliable for the code, but they are structured in different ways.

- 1. The stack store values in order (0,1,2,3,4,5) and removes the value in the opposite order (5,4,3,2,1,0). This concept is called ***last in, first out***.

Adding data  is called ***pushing onto the stack*** and removing data is called ***popping off the stack***.

All data stored in the stack must have a known, fixed data size. Data with an unknown size at compile time or mutable size must be stored in the heap instead.

> The Heap

The heap is less organized. When data is placed into this "area", the program request a certain data amount.The memory allocator perform *** allocating on the heap** also called **allocating**:

1. Finds an empty slot located in the heap.
2. Marks as being in use.
3. Return a poiner, which is the original adress for that location

As the heap pointer is defined and know being fixed data, the porgrammer can store the pointer on the stack, but the data must follow the pointer.

- A example presented in the book is of an costumer in a restaurant with friends:

1. When the group arrives, it declares the number of people within the group

2. The host locate an empty space that fits the whole group and leads there.

3. As more people arrives from the group, they can ask where the group is seated.

### Part 2 : Pushing and Accessing data Comparisons

- Pushing data to the stack is faster than allocating on the heap, because the allocator never has to search for a place to store new data, as that location is always at the top of the stack.

- Allocating data in the heap is more work, because the allocator needs to find first a memory location that is big enough to hold the data and peform bookeeping to prepare the next allocation.

- When the code calls a function, the values passed into that function (data pointers on the heap) and the local variables get placed back onto the stack. When the function ends, these values are placed off the stack.

- Accessing data located in the heap is also slower. Modern CPU's are fasted and use less memory to perform these tasks.

## Ownership Rules

1. Each value in Rust has a owner.

2. There can only be one owner at a time.

3. When the owner goes out of scope, the value will be dropped

### Variable Scope

- The code will no longer provide the main function in examples.

``` rust

fn main() {
// variable s is not valid, as it's not declared 
    let s = "hello"; // variable s is declared, being a literal string
    // code for this variable
}
    // code is over
```

- When s comes to the scope, it's a valid
- The variable s is valid until is out of scope.

### The String Type

The String variable is of unknown size, muttable and user-input friendly.

``` rust
fn main(){
let s = String::from("hello");
}
```

- The **::** allow us to namespce this from function under the String rather than other naming convnetion.

- This type of string can be mutated:

``` rust
fn main(){
    let mut s = String::from("hello");
    s.push_str(", world!"); // push_str pushes a literal to a String
    println!("{}" , s) // This will print 'hello, world!'
}

```

- String literal is different than String. String alone is not hardcoded allowing to be mutable. growable piece of text.

> Memory and Allocation

- With the String type, in order to support a mutable, growable piece of text, we need to allocate an amount of memory on the heap, unknown at compile time, to hold the contents. This means:

1. The memory must be requested from the memory allocator at runtime.
2. We need a way of returning this memory to the allocator when we’re done with our String.

- But, Rust takes a different path: the memory is automatically returned once the variable that owns it goes out of scope. Here’s a version of our scope example from Listing 4-1 using a String instead of a string literal:

``` rust
fn main(){
    // variable s is not valid
    let s = String::from("hello"); // variable s is declared 
    // code about variable s

}
// varaible s is not longer valid or under scope.
```

- There is a natural point at which we can return the memory our String needs to the allocator: when s goes out of scope. When a variable goes out of scope, Rust calls a special function for us. This function is called **drop**, and it’s where the author of String can put the code to return the memory. Rust calls drop automatically at the closing curly bracket.

> Variables and Data Interacting with Move

- Multiple variables can interact with the same data in Rust.

``` rust
// Int version
fn main(){
    let x = 5;
    let y = x;
// x is 5
// y has the same value as x
}
```

``` rust
// String version
fn main(){
    let s1 = String::from("hello");
    let s2 = s1;
// s1 output is hello
// s2 copies the output from s1
}
```

- In the String version, the pointer, length, and capacity is stored in the stack. Only the pointer is copied not the data.

This code will generate an error, but it's useful to understand this example:

``` rust
fn main(){
    let s1 = String::from("hello");
    let s2 = s1;
    println!("{}, world!", s1);
}
```

Rust will generates this error

``` bash

$ cargo run
   Compiling ownership v0.1.0 (file:///projects/ownership)
error[E0382]: borrow of moved value: `s1`
 --> src/main.rs:5:28
  |
2 |     let s1 = String::from("hello");
  |         -- move occurs because `s1` has type `String`, which does not implement the `Copy` trait
3 |     let s2 = s1;
  |              -- value moved here
4 |
5 |     println!("{}, world!", s1);
  |                            ^^ value borrowed here after move
  |
  = note: this error originates in the macro `$crate::format_args_nl` which comes from the expansion of the macro `println` (in Nightly builds, run with -Z macro-backtrace for more info)

For more information about this error, try `rustc --explain E0382`.
error: could not compile `ownership` due to previous error
```

> Variable and Data Interacting with Clone

If we do want to deeply copy the heap data of the String, not just the stack data, we can use a common method called **clone**. This is an example of this method

``` rust
fn main(){
    let s1 = String::from("hello");
    let s2 = s1.clone();
    println!("s1 = {} , s2 = {}", s1,s2);
}
```

The code generates this output

``` bash
s1 = hellom s2 = hello
```

> Stack-Only Data: Copy

Look at the code bwllow:

``` rust
fn main() {
    let x = 5;
    let y = x;

    println!("x = {}, y = {}", x, y);
}

```

This will be the output:

``` bash
x = 5, y = 5
```

- The reason is that types such as integers that have a known size at compile time are stored entirely on the stack, so copies of the actual values are quick to make. That means there’s no reason we would want to prevent x from being valid after we create the variable y.
- So, what types implement the Copy trait? You can check the documentation for the given type to be sure, but as a general rule, any group of simple scalar values can implement Copy, and nothing that requires allocation or is some form of resource can implement Copy. Here are some of the types that implement Copy:

1. All the integer types, such as u32.
2. The Boolean type, bool, with values true and false.
3. All the floating-point types, such as f64.
4. The character type, char.
5. Tuples, if they only contain types that also implement Copy. For example, (i32, i32)
6. Implements Copy, but (i32, String) does not.

> Ownership and Functions

- The method of passing down a value to other functions are similar to those when assign a value to a variable. Look at the code bellow:

``` rust
fn main() {
    let s = String::from("hello");  // s comes into scope

    takes_ownership(s);             // s's value moves into the function...
                                    // ... and so is no longer valid here

    let x = 5;                      // x comes into scope

    makes_copy(x);                  // x would move into the function,
                                    // but i32 is Copy, so it's okay to still
                                    // use x afterward

} // Here, x goes out of scope, then s. But because s's value was moved, nothing
  // special happens.

fn takes_ownership(some_string: String) { // some_string comes into scope
    println!("{}", some_string);
} // Here, some_string goes out of scope and `drop` is called. The backing
  // memory is freed.

fn makes_copy(some_integer: i32) { // some_integer comes into scope
    println!("{}", some_integer);
} // Here, some_integer goes out of scope. Nothing special happens.
```

**Result :**

``` bash
hello
5
```

> Return Value and Scope

- Return value can also transfer ownership. The code bellow explains it:

``` rust 
fn main() {
    let s1 = gives_ownership();        

    let s2 = String::from("hello");     
    let s3 = takes_and_gives_back(s2); 
}
// Main Function Comments
// 1. It gives ownership and moves its return value into s1
// 2. Variable s2 comes into scope
// 3. Variable is moved into this variable, which is also moved its return value into variable s3.
// 4. Here, s3 goes out of scope and its droped. Both variable s1 and s2 is droped and out away of scope.

fn gives_ownership() -> String {             

    let some_string = String::from("yours"); 

    some_string                             
}

// Give_ownership Notes
// 1. Give_onwership will move its return value into the value. And then it calls the value.
// 2. The variable "some-string" comes into scope.
// 3. The variable "some_string" is returned and moves out to the calling function.


fn takes_and_gives_back(a_string: String) -> String { 
    a_string 
}

// Takes and Gives back Notes
// 1. This function takes and return one
// 2. the variable "a string" comes into scope
// 3. the variable "a string" is returnhed and moves out to the calling function.
```

> Transferring ownership of return values

- The onwership of a variable follows the same pathern every time:

1. Assign a value to a variable
2. Move a variable to another function as an argument
3. Move a variable to another function as a return value

``` rust
fn main() {
// This function takes ownership of the heap data
    let s1 = String::from("hello");

    let (s2, len) = calculate_length(s1);

    println!("The length of '{}' is {}.", s2, len);
}
// This function takes ownership of a String and returns the length of the String

fn calculate_length(s: String) -> (String, usize) {
    let length = s.len(); // len() returns the length of a String.

    (s, length)
}
```
