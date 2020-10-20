import std.stdio;

void main() {
   // example 1
   auto f1 = function bool (string s, char c) {
      return s[0] == c;
   };
   // example 2
   auto f2 = (string s, char c) {
      return s[0] == c;
   };
   // example 3
   auto f3 = function bool (string s, char c) => s[0] == c;
   // example 4
   auto f4 = (string s, char c) => s[0] == c;
   // print
   auto b1 = f1("June", 'J');
   auto b2 = f2("June", 'J');
   auto b3 = f3("June", 'J');
   auto b4 = f4("June", 'J');
   writeln(b1 && b2 && b3 && b4);
}
