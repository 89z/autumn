let c = require('crypto');
let [plain, pass] = ['January February', 'PassPassPassPass'];
let cipher = c.createCipheriv('aes-128-ecb', pass, null).setAutoPadding(false);
let ct = cipher.update(plain, 'utf8', 'base64') + cipher.final('base64');
console.log(ct == 'LxPTC0HIdnxnCj54rzkEjA==');
