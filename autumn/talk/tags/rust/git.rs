#![allow(unused)]

use git2::{
   DiffFormat,
   Repository
};

fn init(s: &str) -> Result<git2::Repository, git2::Error> {
   Repository::init(s)
}

fn clone(url_s: &str, path_s: &str) -> Result<git2::Repository, git2::Error> {
   Repository::clone(url_s, path_s)
}

fn main() -> Result<(), git2::Error> {
   let repo_o = Repository::open(r"D:\Desktop\rosso")?;
   let diff_o = repo_o.diff_index_to_workdir(None, None)?;
   diff_o.print(DiffFormat::Patch, |_delta, _hunk, line| {
      let s = String::from_utf8_lossy(line.content());
      if line.origin() == '-' {
         print!("\x1b[31m{}\x1b[m", s);
      } else {
         print!("{}", s);
      }
      true
   });
   Ok(())
}
