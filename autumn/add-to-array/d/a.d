import std.stdio;

void main() {
   auto a = [0, 10];
   // example 1
   a ~= 20;
   // example 2
   a ~= [30, 40];
   // print
   a.writeln;
}
