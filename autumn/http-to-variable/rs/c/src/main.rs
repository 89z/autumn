fn main() -> Result<(), attohttpc::Error> {
   let s = "https://speedtest.lax.hivelocity.net";
   let o = attohttpc::get(s).send()?;
   let s = o.text()?;
   print!("{}", s);
   Ok(())
}
