import std.stdio;

void main() {
   // example 1
   byte n1 = 0x7F;
   // example 2
   ubyte n2 = 0xFF;
   // example 3
   short n3 = 0x7FFF;
   // example 4
   ushort n4 = 0xFFFF;
   // example 5
   int n5 = 0x7FFF_FFFF;
   // example 6
   uint n6 = 0xFFFF_FFFF;
   // example 7
   long n7 = 0x7FFF_FFFF_FFFF_FFFF;
   // example 8
   ulong n8 = 0xFFFF_FFFF_FFFF_FFFF;
   // print
   [n1, n2, n3, n4, n5, n6, n7, n8].writeln;
}
