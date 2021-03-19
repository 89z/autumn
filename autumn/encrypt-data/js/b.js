import { Aes } from 'http://deno.land/x/crypto/aes.ts';
import { Cbc, Padding } from 'http://deno.land/x/crypto/block-modes.ts';
let te = new TextEncoder;

let [plain, pass, iv] = [
   te.encode('January February'),
   te.encode('PassPassPassPass'),
   te.encode('IvIvIvIvIvIvIvIv')
];

let cipher = new Cbc(Aes, pass, iv, Padding.PKCS7);
let ct = btoa(String.fromCharCode(...cipher.encrypt(plain)));
console.log(ct == 'x2B4PSY4exW+U7x/jmOkSJmDsB0IgpsLeHSICOQnPF8=');
