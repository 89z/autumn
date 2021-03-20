import nimcrypto
var aliceKey = "Alice Key"
var aliceData = "Alice hidden secret"
var ectx, dctx: ECB[aes256]
var key = newString(aes256.sizeKey)
var plainText = newString(aes256.sizeBlock * 2)
var encText = newString(aes256.sizeBlock * 2)
var decText = newString(aes256.sizeBlock * 2)
copyMem(addr plainText[0], addr aliceData[0], len(aliceData))
copyMem(addr key[0], addr aliceKey[0], len(aliceKey))
ectx.init(key)
ectx.encrypt(plainText, encText)
ectx.clear()
dctx.init(key)
dctx.decrypt(encText, decText)
dctx.clear()
echo "PLAIN TEXT: ", $plainText
echo "ENCODED TEXT: ", $encText
echo "DECODED TEXT: ", $decText
assert(equalMem(addr plainText[0], addr decText[0], len(plainText)))
