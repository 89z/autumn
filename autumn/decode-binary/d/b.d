import io = std.stdio;
import std.base64: Base64;

void main() {
   auto s = "CgsMDQ==";
   auto a = Base64.decode(s);
   io.writeln(a == [10, 11, 12, 13]);
}
