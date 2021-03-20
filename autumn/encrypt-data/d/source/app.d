import botan.all;
import botan.filters.b64_filt;
import botan.filters.hex_filt;
import std.stdio;

void main() {
   auto pipe = Pipe(new HexEncoder);
   pipe.processMsg("PassPassPassPass");
   pipe = Pipe(
      getCipher("AES-128/ECB", pipe.toString.OctetString, ENCRYPTION),
      new Base64Encoder
   );
   pipe.processMsg("January February");
   writeln(pipe.toString == "LxPTC0HIdnxnCj54rzkEjA==");
}
