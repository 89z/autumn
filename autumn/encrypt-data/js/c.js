let crypto = require('crypto');

let [plain, pass, iv] = [
   'January February', 'PassPassPassPass', 'IvIvIvIvIvIvIvIv'
];

let cipher = crypto.createCipheriv('aes-128-cbc', pass, iv);
let ct = cipher.update(plain, 'utf8', 'base64') + cipher.final('base64');
console.log(ct == 'x2B4PSY4exW+U7x/jmOkSJmDsB0IgpsLeHSICOQnPF8=');
