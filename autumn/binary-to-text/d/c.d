import io = std.stdio;
import std.base64: Base64;

void main() {
   ubyte[] a = [10, 11, 12, 13];
   string s = Base64.encode(a);
   io.writeln(s == "CgsMDQ==");
}
