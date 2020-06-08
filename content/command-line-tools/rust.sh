# example 1
rustc a.rs
# example 2
cargo new sunday
cd sunday
cat > Cargo.toml <<eof
[package]
name = "sunday"
version = "1.0.0"
edition = "2018"
[dependencies]
regex-automata = ""
eof
cargo build
