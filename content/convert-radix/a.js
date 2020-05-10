let s_dig = '-0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ_abcdefghijklmnopqrstuvwxyz';
// example 1
function r64_encode(n_in) {
   let s_out = '';
   while (n_in > 0) {
      s_out = s_dig[n_in & 63] + s_out;
      n_in >>= 6;
   }
   return s_out;
}
// example 2
function r64_decode(s_in) {
   let n_out = 0;
   for (let s_chr of s_in) {
      n_out = n_out << 6 | s_dig.indexOf(s_chr);
   }
   return n_out;
}
// print
let n1 = Math.trunc(new Date / 1000);
let s1 = r64_encode(n1);
let n2 = r64_decode(s1);
console.log(n1, s1, n2 == n1);
