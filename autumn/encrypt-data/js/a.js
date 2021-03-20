let [te, cs] = [new TextEncoder, crypto.subtle];

let [plain, keyData, iv] = [
   te.encode('January February'),
   te.encode('KDKDKDKDKDKDKDKD'),
   te.encode('IVIVIVIVIVIVIVIV')
];

async function main() {
   let key = await cs.importKey('raw', keyData, 'AES-CBC', true, ['encrypt']);
   let cipher = await cs.encrypt({name: 'AES-CBC', iv}, key, plain);
   let ct = btoa(String.fromCharCode(...new Uint8Array(cipher)));
   console.log(ct == 'BvfnZp4jmCaveE6kefhumpZ0raWX9GDojfPasgSwLTM=');
}

main();
