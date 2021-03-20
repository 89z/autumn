import { Aes } from 'http://deno.land/x/crypto/aes.ts';
import { Ecb } from 'http://deno.land/x/crypto/block-modes.ts';
let te = new TextEncoder;

let [plain, key] = [
   te.encode('January February'), te.encode('KDKDKDKDKDKDKDKD')
];

let cipher = new Ecb(Aes, key);
let ct = btoa(String.fromCharCode(...cipher.encrypt(plain)));
console.log(ct == 'hr0e+xH2oi0mYHMmdGQCnQ==');
