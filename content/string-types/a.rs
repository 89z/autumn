fn main() {
   // example 1
   let s1 = "Sunday";
   // example 2
   let s2 = String::from("Sunday");
   // example 3
   let s3 = "Sunday
Monday";
   // example 4
   let s4 = r#"Sun "Mon" Tue"#;
   // print
   println!("{}{}{}{}", s1, s2, s3, s4);
}
