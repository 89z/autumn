import io = std.stdio;
import std.base64;

void main() {
   ubyte[] data = [0x14, 0xfb, 0x9c, 0x03, 0xd9, 0x7e];
   const(char)[] encoded = Base64.encode(data);
   assert(encoded == "FPucA9l+");
   ubyte[] decoded = Base64.decode("FPucA9l+");
   assert(decoded == [0x14, 0xfb, 0x9c, 0x03, 0xd9, 0x7e]);
}
