import std.stdio;

auto startsWith(string s, char c) {
   return s[0] == c;
}

void main() {
   auto b = startsWith("May", 'M');
   writeln(b);
}
