import std.stdio;

void main() {
   char c = '\x43';
   int n;

   switch (c) {
   case 'A':
      n = 0x41;
      break;
   case 'B', 'b':
      n = 0x42;
      break;
   default:
      n = cast(int) c;
   }

   writeln(n == 0x43);
}
