import botan.all;
import botan.filters.b64_filt;
import botan.filters.hex_filt;
import std.stdio;

void main() {
   auto p = Pipe(new HexEncoder);
   p.processMsg("PassPassPassPass");
   auto key = p.toString.SymmetricKey;
   p.processMsg("IvIvIvIvIvIvIvIv");
   auto iv = p.toString.InitializationVector;
   p = Pipe(getCipher("AES-128/CBC", key, iv, ENCRYPTION), new Base64Encoder);
   // The nonce is fresh for this message
   p.processMsg("January February");
   writeln(p.toString);
}
