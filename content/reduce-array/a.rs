fn main() {
   let arr = vec![1, 2, 3, 4, 5, 6, 7, 8, 9];
   let sum = arr.iter().fold(0i32, |a, &b| a + b);
   let product = arr.iter().fold(1i32, |a, &b| a * b);
   println!("the sum is {} and the product is {}", sum, product);
}
