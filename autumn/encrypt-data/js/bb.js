import { Aes } from 'http://deno.land/x/crypto/aes.ts';
import { Ecb } from 'http://deno.land/x/crypto/block-modes.ts';
let te = new TextEncoder;

let [plain, pass] = [
   te.encode('January February'), te.encode('PassPassPassPass')
];

let cipher = new Ecb(Aes, pass);
let ct = btoa(String.fromCharCode(...cipher.encrypt(plain)));
console.log(ct == 'LxPTC0HIdnxnCj54rzkEjA==');
