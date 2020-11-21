import std.stdio;

void main() {
   int n = 0x63 - 0x20;
   char c;

   switch (n) {
   case 0x41:
      c = 'A';
      break;
   case 0x42, 0x62:
      c = 'B';
      break;
   default:
      c = cast(char) n;
   }

   writeln(c == 'C');
}
