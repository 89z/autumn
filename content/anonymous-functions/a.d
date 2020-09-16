import std.stdio;
// example 1
bool f(int n, int n2) {
   return n > n2;
}
// example 2
auto f2(int n, int n2) {
   return n > n2;
}
// example 3
auto f3 = function bool (int n, int n2) {
   return n > n2;
};
// example 4
auto f4 = (int n, int n2) {
   return n > n2;
};
// example 5
auto f5 = function bool (int n, int n2) => n > n2;
// example 6
auto f6 = (int n, int n2) => n > n2;
// print
void main() {
   auto b = f(9, 8);
   auto b2 = f2(9, 8);
   auto b3 = f3(9, 8);
   auto b4 = f4(9, 8);
   auto b5 = f5(9, 8);
   auto b6 = f6(9, 8);
   writeln(b && b2 && b3 && b4 && b5 && b6);
}
