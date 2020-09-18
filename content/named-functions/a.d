import std.stdio;
// example 1
bool f1(int n1, int n2) {
   return n1 > n2;
}
// example 2
auto f2(int n1, int n2) {
   return n1 > n2;
}
// print
void main() {
   auto b = f(9, 8);
   auto b2 = f2(9, 8);
   writeln(b && b2);
}
