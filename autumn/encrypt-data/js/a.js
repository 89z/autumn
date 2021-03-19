let [te, cs] = [new TextEncoder, crypto.subtle];

let [plain, pass, iv] = [
   te.encode('January February'),
   te.encode('PassPassPassPass'),
   te.encode('IvIvIvIvIvIvIvIv')
];

async function main() {
   let key = await cs.importKey('raw', pass, 'AES-CBC', true, ['encrypt']);
   let cipher = await cs.encrypt({name: 'AES-CBC', iv}, key, plain);
   let ct = btoa(String.fromCharCode(...new Uint8Array(cipher)));
   console.log(ct == 'x2B4PSY4exW+U7x/jmOkSJmDsB0IgpsLeHSICOQnPF8=');
}

main();
