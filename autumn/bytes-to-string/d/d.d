import std.digest, std.stdio;

void main() {
   { // example 1
      ubyte[] a = [10, 11, 12, 13];
      string s = a.toHexString;
      writeln(s == "0A0B0C0D");
   }
   { // example 2
      string s = "\n\v\f\r";
      string t = toHexString(cast(ubyte[])s);
      writeln(t == "0A0B0C0D");
   }
}
