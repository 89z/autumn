# example 1
rustc a.rs
# example 2
cargo new sunday
Set-Location sunday
'
[package]
name = "sunday"
version = "1.0.0"
edition = "2018"
[dependencies]
curl = ""
' > Cargo.toml
cargo build
