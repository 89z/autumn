let c = require('crypto');
let [plain, key] = ['January February', 'KDKDKDKDKDKDKDKD'];
let cipher = c.createCipheriv('aes-128-ecb', key, null).setAutoPadding(false);
let ct = cipher.update(plain, 'utf8', 'base64') + cipher.final('base64');
console.log(ct == 'hr0e+xH2oi0mYHMmdGQCnQ==');
