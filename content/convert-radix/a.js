let s_dig = '0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ-_';
// example 1
function r64_encode(n_in) {
   let s_out = '';
   for (let n_sh = n_in; n_sh > 0; n_sh >>= 6) {
      let n_key = n_sh & 63;
      s_out = s_dig[n_key] + s_out;
   }
   return s_out;
}
// example 2
function r64_decode(s_in) {
   let n_out = 0;
   let n_key = s_in.length - 1;
   for (let n_sh = 0; n_sh < 36; n_sh += 6) {
      let s_val = s_in[n_key];
      n_out |= s_dig.indexOf(s_val) << n_sh;
      n_key -= 1
   }
   return n_out;
}
// print
let n1 = Math.trunc(new Date / 1000);
let s1 = r64_encode(n1);
let n2 = r64_decode(s1);
console.log(n1, s1, n2 == n1);
