use filetime::FileTime;

fn main() {
   let o = FileTime::from_unix_time(400_000_000, 0);
   let e = filetime::set_file_mtime("Cargo.toml", o);
   println!("{:?}", e);
}
