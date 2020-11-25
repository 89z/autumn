use {
   csv::Reader,
   std::collections::HashMap
};

fn main() -> Result<(), csv::Error> {
   let mut o = Reader::from_path("a.csv")?;
   for e in o.deserialize() {
      let m: HashMap<String, String> = e?;
      println!("{:?}", m);
   }
   Ok(())
}
