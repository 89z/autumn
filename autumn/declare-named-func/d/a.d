import std.stdio;

int add(int n, int n2) {
   return n + n2;
}

void main() {
   auto n = 7;
   auto n2 = n.add(1).add(1);
   writeln(n2 == 9);
}
