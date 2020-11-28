import std.stdio;

auto add(int n, int n2) {
   return n + n2;
}

void main() {
   auto n = add(8, 1);
   writeln(n == 9);
}
