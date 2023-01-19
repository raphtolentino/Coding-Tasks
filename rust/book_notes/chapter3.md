# Chapter 3
> Variables and Mutability 

- Rust variables are immutable.
-** Once its bound to a name, you cannot change its value **

### Example 1.0


``` rust
fn main() {
    let x = 5;
    println!("The value of x is {x}");
    x = 6;
    println!("The value of x is {x}");
}

``` 
It prints this error code:

``` 
error[E0384]: cannot assign twice to immutable variable `x`
 --> src\main.rs:4:5
  |
2 |     let x = 5;
  |         -
  |         |
  |         first assignment to `x`
  |         help: consider making this binding mutable: `mut x`
3 |     println!("The value of x is {x}");
4 |     x = 6;
  |     ^^^^^ cannot assign twice to immutable variable

  ```

- This error showed because, I cannot assign two immutable variables "x"

