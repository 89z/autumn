// example 1
num f(num d, num e) {
   return d + e;
}

// example 2
num g(num d, num e) => d + e;

// print
void main() {
   print(f(4, 5) == 9);
   print(g(4, 5) == 9);
}
