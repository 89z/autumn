use std::collections::VecDeque;

fn main() {
   let mut a = VecDeque::new();
   a.push_front(10);
   a.push_front(11);
   println!("{:?}", a);
}
