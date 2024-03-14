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