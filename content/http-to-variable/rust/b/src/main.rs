use minreq;

fn main() -> Result<(), minreq::Error> {
   let s_in = "https://speedtest.lax.hivelocity.net";
   let o = minreq::get(s_in).send()?;
   let s_out = o.as_str()?;
   print!("{}", s_out);
   Ok(())
}
