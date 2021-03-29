let crypto = require('crypto');

let [plain, key, iv] = [
   'north east south', 'KDKDKDKDKDKDKDKD', 'IVIVIVIVIVIVIVIV'
];

let cipher = crypto.createCipheriv('aes-128-cbc', key, iv);
let ct = cipher.update(plain, 'utf8', 'base64') + cipher.final('base64');
console.log(ct == 'gGPdBt1qCJCbHye4HX1E9ZUhYmjPsvEi7AsIs3XOReg=');
