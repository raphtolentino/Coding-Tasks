# Chapter 4 - Understanding Ownership

> Chapter 4.1 - What is Ownership?

Ownership refers to a rule set that dictates for Rust manages memory.

- 1. Some program languages has a garbage collection program that runs together with the compiler that regurarly looks for non-longer-used memory as that the program is running.

- 2. Some program languages the programmer has to explicitly allocate the variables to free memory.

- 3. Rust is a bit different, it uses a ownership system with a set of rules that checks alongside of the company. If any rules are violeted the program will not run.

## The Stack and the Heap

### The Stack

- Many different programing languages, this concept of stack and heap is not as relevant, but in Rust it's affects how the language behaves.

- Both the stack and the heap are part of memory avaliable for the code, but they are structured in different ways.

- 1. The stack store values in order (0,1,2,3,4,5) and removes the value in the opposite order (5,4,3,2,1,0). This concept is called ***last in, first out***.

Adding data  is called ***pushing onto the stack*** and removing data is called ***popping off the stack***.

All data stored in the stack must have a known, fixed data size. Data with an unknown size at compile time or mutable size must be stored in the heap instead.

### The Heap

- The heap is less organized. When data 