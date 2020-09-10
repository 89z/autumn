# example 1
rustc a.rs
# example 2
cargo new may
# example 3
'
[package]
name = "may"
version = "1.0.0"
edition = "2018"
[dependencies]
curl = ""
' > Cargo.toml
# example 4
cargo run
# example 5
$env:RUSTFLAGS = '-C link-arg=-s'
cargo build --release
