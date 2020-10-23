use h2::client;

use http::{Request, Method};
use std::error::Error;
use tokio::net::TcpStream;

#[tokio::main]
pub async fn main() -> Result<(), Box<dyn Error>> {
    // Establish TCP connection to the server.
    let tcp = TcpStream::connect("127.0.0.1:5928").await?;
    let (h2, connection) = client::handshake(tcp).await?;
    tokio::spawn(async move {
        connection.await.unwrap();
    });

    let mut h2 = h2.ready().await?;
    // Prepare the HTTP request to send to the server.
    let request = Request::builder()
                    .method(Method::GET)
                    .uri("https://www.example.com/")
                    .body(())
                    .unwrap();

    // Send the request. The second tuple item allows the caller
    // to stream a request body.
    let (response, _) = h2.send_request(request, true).unwrap();

    let (head, mut body) = response.await?.into_parts();

    println!("Received response: {:?}", head);

    // The `flow_control` handle allows the caller to manage
    // flow control.
    //
    // Whenever data is received, the caller is responsible for
    // releasing capacity back to the server once it has freed
    // the data from memory.
    let mut flow_control = body.flow_control().clone();

    while let Some(chunk) = body.data().await {
        let chunk = chunk?;
        println!("RX: {:?}", chunk);

        // Let the server send more data.
        let _ = flow_control.release_capacity(chunk.len());
    }

    Ok(())
}
