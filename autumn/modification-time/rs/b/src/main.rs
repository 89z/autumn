use {
   filetime::FileTime,
   std::io
};

fn main() -> io::Result<()> {
   let o = FileTime::from_unix_time(400_000_000, 0);
   filetime::set_file_mtime("Cargo.toml", o)
}
