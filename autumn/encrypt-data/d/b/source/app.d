import botan.all;
import botan.filters.b64_filt;
import botan.filters.hex_filt;
import std.stdio;

void main() {
   auto p = Pipe(new HexEncoder);
   p.processMsg("PassPassPassPass");
   p = Pipe(
      getCipher("AES-128/ECB", p.toString.SymmetricKey, ENCRYPTION),
      new Base64Encoder
   );
   p.processMsg("January February");
   writeln(p.toString == "LxPTC0HIdnxnCj54rzkEjA==");
}
