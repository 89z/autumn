fn main() {
    let x = "été chaud";
    
    let y = "chaud";
    let i = x.find(y);
    println!("{:?}", i);
    
    let y = "froid";
    let i = x.find(y);
    println!("{:?}", i);
}
