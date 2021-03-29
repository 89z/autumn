import botan.all, botan.filters.b64_filt, std.stdio;

void main() {
   auto cipher = getCipher("AES-128/ECB", ENCRYPTION);
   cipher.setKey(SymmetricKey("4B444B444B444B444B444B444B444B44"));
   auto pipe = Pipe(cipher, new Base64Encoder);
   pipe.processMsg("north east south");
   writeln(pipe.toString == "4HbK4aRsiiLXEtT99V2Xgg==");
}
