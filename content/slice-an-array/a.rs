fn main() {
   let a1 = vec![10, 11, 12];
   // example 1
   let n1 = a1[0];
   // example 2
   let n2 = a1[a1.len() - 1];
   // example 3
   let n3 = a1.get(0);
   // print
   dbg!(n1, n2, n3);
}
