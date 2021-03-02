import std.stdio;

void main() {
   // example 1
   byte d = 0x7F;
   // example 2
   ubyte e = 0xFF;
   // example 3
   short f = 0x7FFF;
   // example 4
   ushort g = 0xFFFF;
   // example 5
   int h = 0x7FFF_FFFF;
   // example 6
   uint i = 0xFFFF_FFFF;
   // example 7
   long j = 0x7FFF_FFFF_FFFF_FFFF;
   // example 8
   ulong k = 0xFFFF_FFFF_FFFF_FFFF;
   // print
   [d, e, f, g, h, i, j, k].writeln;
}
