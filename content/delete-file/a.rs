#![allow(unused)]
use std::fs;

fn main() -> std::io::Result<()> {
    fs::remove_file("a.txt")?;
    Ok(())
}
