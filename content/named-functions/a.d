import std.stdio;
// example 1
bool f1(int n, int n1) {
   return n > n1;
}
// example 2
auto f2(int n, int n2) {
   return n > n2;
}
// print
void main() {
   auto b1 = f1(9, 8);
   auto b2 = f2(9, 8);
   writeln(b1 && b2);
}
