use {
   http_req::request,
   std::error::Error,
   std::io,
   std::io::Write
};

fn main() -> Result<(), Box<dyn Error>> {
   let mut a = Vec::new();
   request::get("https://speedtest.lax.hivelocity.net", &mut a)?;
   io::stdout().write(&a)?;
   Ok(())
}
