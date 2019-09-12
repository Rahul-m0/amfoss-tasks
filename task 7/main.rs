extern crate regex;
use regex::Regex;
use std::io;

fn main() {
    let re = Regex::new(r"^[a-zA-Z0-9]+@[a-z][.]?[a-z]*").unwrap();
    let emailid = String::new();
    println!("Enter your email:", );
    let mut emailid = String::new();
    io::stdin().read_line(&mut emailid).expect("Error Please try again later"); 
    if re.is_match(&mut emailid) == true 
     {
        println!("What you entered is a valid email format", )
    }
    else {
        println!("What you entered is an invalid email format", )
    }
}
