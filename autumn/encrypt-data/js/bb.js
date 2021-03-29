import { Aes } from 'http://deno.land/x/crypto/aes.ts';
import { Ecb } from 'http://deno.land/x/crypto/block-modes.ts';
let te = new TextEncoder;

let [plain, key] = [
   te.encode('north east south'), te.encode('KDKDKDKDKDKDKDKD')
];

let cipher = new Ecb(Aes, key);
let ct = btoa(String.fromCharCode(...cipher.encrypt(plain)));
console.log(ct == '4HbK4aRsiiLXEtT99V2Xgg==');
