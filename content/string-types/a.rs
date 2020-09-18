fn main() {
   // example A
   let s_a = "May";
   // example B
   let s_b = String::from("May");
   // example C
   let s_c = "May
June";
   // example D
   let s_d = r#"May "June" July"#;
   // print
   println!("{},{},{},{}", s_a, s_b, s_c, s_d);
}
