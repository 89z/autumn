import botan.all;
import botan.filters.b64_filt;
import std.stdio;

void main() {
   auto cipher = getCipher("AES-128/CBC", ENCRYPTION);
   cipher.setKey(SymmetricKey("4B444B444B444B444B444B444B444B44"));
   cipher.setIv(InitializationVector("49564956495649564956495649564956"));
   auto pipe = Pipe(cipher, new Base64Encoder);
   pipe.processMsg("January February");
   writeln(pipe.toString == "BvfnZp4jmCaveE6kefhumpZ0raWX9GDojfPasgSwLTM=");
}
