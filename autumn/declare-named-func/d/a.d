import std.stdio;

int f(int d, int e) {
   return d + e;
}

void main() {
   auto n = f(4, 5);
   writeln(n == 9);
}
