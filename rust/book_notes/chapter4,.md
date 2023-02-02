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

image.png

