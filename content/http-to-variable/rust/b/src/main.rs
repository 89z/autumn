use {
   http_req::request,
   std::error::Error
};

fn main() -> Result<(), Box<dyn Error>> {
   let s = "https://speedtest.lax.hivelocity.net";
   let mut a = Vec::new();
   request::get(s, &mut a)?;
   let s = String::from_utf8(a)?;
   print!("{}", s);
   Ok(())
}
