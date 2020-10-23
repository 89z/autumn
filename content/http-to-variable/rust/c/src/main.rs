extern crate solicit;
use solicit::http::client::CleartextConnector;
use solicit::client::SimpleClient;
use std::str;

fn main() {
  // Connect to an HTTP/2 aware server
  let connector = CleartextConnector::new("http2bin.org");
  let mut client = SimpleClient::with_connector(connector).unwrap();
  let response = client.get(b"/get", &[]).unwrap();
  assert_eq!(response.stream_id, 1);
  assert_eq!(response.status_code().unwrap(), 200);
  // Dump the headers and the response body to stdout.
  // They are returned as raw bytes for the user to do as they please.
  // (Note: in general directly decoding assuming a utf8 encoding might not
  // always work -- this is meant as a simple example that shows that the
  // response is well formed.)
  for header in response.headers.iter() {
    println!("{}: {}",
        str::from_utf8(header.name()).unwrap(),
        str::from_utf8(header.value()).unwrap());
  }
  println!("{}", str::from_utf8(&response.body).unwrap());
  // We can issue more requests after reading this one...
  // These calls block until the request itself is sent, but do not wait
  // for a response.
  let req_id1 = client.request(b"GET", b"/get?hi=hello", &[], None).unwrap();
  let req_id2 = client.request(b"GET", b"/asdf", &[], None).unwrap();
  // Now we get a response for both requests... This does block.
  let (resp1, resp2) = (
      client.get_response(req_id1).unwrap(),
      client.get_response(req_id2).unwrap(),
  );
  assert_eq!(resp1.status_code().unwrap(), 200);
  assert_eq!(resp2.status_code().unwrap(), 404);
}
