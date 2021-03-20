import { Aes } from 'http://deno.land/x/crypto/aes.ts';
import { Cbc, Padding } from 'http://deno.land/x/crypto/block-modes.ts';
let te = new TextEncoder;

let [plain, key, iv] = [
   te.encode('January February'),
   te.encode('KDKDKDKDKDKDKDKD'),
   te.encode('IVIVIVIVIVIVIVIV')
];

let cipher = new Cbc(Aes, key, iv, Padding.PKCS7);
let ct = btoa(String.fromCharCode(...cipher.encrypt(plain)));
console.log(ct == 'BvfnZp4jmCaveE6kefhumpZ0raWX9GDojfPasgSwLTM=');
