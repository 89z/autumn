import std.stdio;

void main() {
   // example 1
   string s = "📗";
   auto n1 = s.length;
   // example 2
   wstring s2 = "📗";
   auto n2 = s2.length;
   // example 3
   dstring s3 = "📗";
   auto n3 = s3.length;
   // print
   writeln(n1 == 4 && n2 == 2 && n3 == 1);
}
