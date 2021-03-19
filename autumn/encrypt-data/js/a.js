let [e, s] = [new TextEncoder, crypto.subtle];

let [plain, pass, iv] = [
   e.encode('January February'),
   e.encode('PassPassPassPass'),
   e.encode('IvIvIvIvIvIvIvIv')
];

async function main() {
   let key = await s.importKey('raw', pass, 'AES-CBC', true, ['encrypt']);
   let cipher = await s.encrypt({name: 'AES-CBC', iv}, key, plain);
   let ct = btoa(String.fromCharCode(...new Uint8Array(cipher)));
   console.log(ct == 'x2B4PSY4exW+U7x/jmOkSJmDsB0IgpsLeHSICOQnPF8=');
}

main();
