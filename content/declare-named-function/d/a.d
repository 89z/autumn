import std.stdio;
// example 1
bool f1(string s, char c) {
   return s[0] == c;
}
// example 2
auto f2(string s, char c) {
   return s[0] == c;
}
// print
void main() {
   auto b1 = f1("June", 'J');
   auto b2 = f2("June", 'J');
   writeln(b1 && b2);
}
