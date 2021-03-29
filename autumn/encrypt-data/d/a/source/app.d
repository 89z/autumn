import botan.all, botan.filters.b64_filt, std.stdio;

void main() {
   auto cipher = getCipher("AES-128/CBC", ENCRYPTION);
   cipher.setKey(SymmetricKey("4B444B444B444B444B444B444B444B44"));
   cipher.setIv(InitializationVector("49564956495649564956495649564956"));
   auto pipe = Pipe(cipher, new Base64Encoder);
   pipe.processMsg("north east south");
   writeln(pipe.toString == "gGPdBt1qCJCbHye4HX1E9ZUhYmjPsvEi7AsIs3XOReg=");
}
