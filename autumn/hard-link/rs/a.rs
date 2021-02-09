use std::fs;

fn main() -> std::io::Result<()> {
    fs::hard_link("a.txt", "b.txt")?; // Hard link a.txt to b.txt
    Ok(())
}
