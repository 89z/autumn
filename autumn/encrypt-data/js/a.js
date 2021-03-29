let [te, cs] = [new TextEncoder, crypto.subtle];

let [plain, keyData, iv] = [
   te.encode('north east south'),
   te.encode('KDKDKDKDKDKDKDKD'),
   te.encode('IVIVIVIVIVIVIVIV')
];

async function main() {
   let key = await cs.importKey('raw', keyData, 'AES-CBC', true, ['encrypt']);
   let cipher = await cs.encrypt({name: 'AES-CBC', iv}, key, plain);
   let ct = btoa(String.fromCharCode(...new Uint8Array(cipher)));
   console.log(ct == 'gGPdBt1qCJCbHye4HX1E9ZUhYmjPsvEi7AsIs3XOReg=');
}

main();
