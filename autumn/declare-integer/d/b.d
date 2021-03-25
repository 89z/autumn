import std.stdio;

void main() {
   { // example 1
      byte n = 0x7F;
      n.writeln;
   }
   { // example 2
      ubyte n = 0xFF;
      n.writeln;
   }
   { // example 3
      short n = 0x7FFF;
      n.writeln;
   }
   { // example 4
      ushort n = 0xFFFF;
      n.writeln;
   }
   { // example 5
      int n = 0x7FFF_FFFF;
      n.writeln;
   }
   { // example 6
      uint n = 0xFFFF_FFFF;
      n.writeln;
   }
   { // example 7
      long n = 0x7FFF_FFFF_FFFF_FFFF;
      n.writeln;
   }
   { // example 8
      ulong n = 0xFFFF_FFFF_FFFF_FFFF;
      n.writeln;
   }
}
