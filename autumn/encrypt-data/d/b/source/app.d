import botan.all;
import botan.filters.b64_filt;
import std.stdio;

void main() {
   auto cipher = getCipher("AES-128/ECB", ENCRYPTION);
   cipher.setKey(SymmetricKey("4B444B444B444B444B444B444B444B44"));
   auto pipe = Pipe(cipher, new Base64Encoder);
   pipe.processMsg("January February");
   writeln(pipe.toString == "hr0e+xH2oi0mYHMmdGQCnQ==");
}
