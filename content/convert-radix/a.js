let s_dig = '0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ-_';
// example 1
function r64_encode(n_in) {
   let s_out = '';
   for (let n_sh = n_in; n_sh > 0; n_sh >>= 6) {
      let n_ind = n_sh & 63;
      s_out += s_dig[n_ind];
   }
   return s_out;
}
let s1 = r64_encode(1234567890);
// example 2
function r64_decode(s_in) {
   let n_out = 0;
   for (let n_sh = 0, n_in = 0; n_sh < 36; n_sh += 6, n_in += 1) {
      let s_chr = s_in[n_in];
      n_out |= s_dig.indexOf(s_chr) << n_sh;
   }
   return n_out;
}
let n1 = r64_decode('ibwB91');
// print
console.log(s1 == 'ibwB91', n1 == 1234567890);
