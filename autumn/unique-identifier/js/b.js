'use strict';

function idDecode(s, n) {
   return new Date(
      Date.UTC(n) + parseInt(s, 36) * 1000
   );
}

let t = idDecode('6dv3d', 2020);
console.log(t);
