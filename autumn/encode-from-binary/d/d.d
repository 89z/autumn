import std.digest, std.stdio;

void main() {
   // example 1
   ubyte[] a = [10, 11, 12, 13];
   string s = toHexString(a);
   // example 2
   string m = "March";
   string t = toHexString(cast(ubyte[])m);
   // print
   writeln(s == "0A0B0C0D" && t == "4D61726368");
}
