use minreq;

fn main() -> Result<(), minreq::Error> {
   let s = "https://speedtest.lax.hivelocity.net";
   let o = minreq::get(s).send()?;
   let s = o.as_str()?;
   print!("{}", s);
   Ok(())
}
