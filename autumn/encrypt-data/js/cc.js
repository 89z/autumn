let c = require('crypto');
let [plain, key] = ['north east south', 'KDKDKDKDKDKDKDKD'];
let cipher = c.createCipheriv('aes-128-ecb', key, null).setAutoPadding(false);
let ct = cipher.update(plain, 'utf8', 'base64') + cipher.final('base64');
console.log(ct == '4HbK4aRsiiLXEtT99V2Xgg==');
