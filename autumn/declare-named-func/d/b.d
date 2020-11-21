import std.stdio;

auto f(int n, int n2) {
   return n > n2;
}

void main() {
   auto b = f(9, 8);
   writeln(b);
}
