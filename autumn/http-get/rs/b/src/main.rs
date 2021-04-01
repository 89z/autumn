use {
   http_req::error,
   http_req::request,
   std::io,
   std::io::Write
};

fn main() -> Result<(), error::Error> {
   let mut a = Vec::new();
   request::get("https://speedtest.lax.hivelocity.net", &mut a)?;
   io::stdout().write(&a)?;
   Ok(())
}
