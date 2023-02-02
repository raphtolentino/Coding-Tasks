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

## Part 2 : Pushing and Accessing data Comparisons

- Pushing data to the stack is faster than allocating on the heap, because the allocator never has to search for a place to store new data, as that location is always at the top of the stack.

- Allocating data in the heap is more work, because the allocator needs to find first a memory location that is big enough to hold the data and peform bookeeping to prepare the next allocation.

- When the code calls a function, the values passed into that function (data pointers on the heap) and the local variables get placed back onto the stack. When the function ends, these values are placed off the stack.

- Accessing data located in the heap is also slower. Modern CPU's are fasted and use less memory to perform these tasks.
