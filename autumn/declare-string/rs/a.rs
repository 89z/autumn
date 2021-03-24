fn main() {
   let mut s;

   // slash
   s = r"south\north";
   println!("{}", s);

   // quote
   s = r#"south"north"#;
   println!("{}", s);

   // newline
   s = r"south
north";
   println!("{}", s);

}
