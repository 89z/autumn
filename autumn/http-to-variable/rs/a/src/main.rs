use minreq;

fn main() -> Result<(), minreq::Error> {
   let s = "https://speedtest.lax.hivelocity.net";
   let r = minreq::get(s).send()?;
   let s = r.as_str()?;
   print!("{}", s);
   Ok(())
}
