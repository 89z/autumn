function Radix64() {
   let s = '-0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ_abcdefghijklmnopqrstuvwxyz';
   this.s_dig = s;
}

Radix64.prototype.encode = function (n_in) {
   let s_out = '';
   do {
      s_out = this.s_dig[n_in % 64] + s_out;
      n_in = Math.trunc(n_in / 64);
   } while (n_in > 0);
   return s_out;
};

Radix64.prototype.decode = function (s_in) {
   let n_out = 0;
   for (let s_chr of s_in) {
      n_out = n_out * 64 + this.s_dig.indexOf(s_chr);
   }
   return n_out;
};

let n = 1577858399;
let o = new Radix64;
// example 1
let s1 = o.encode(n);
let n1 = o.decode(s1);
// example 2
let s2 = o.encode(n - 1);
let n2 = o.decode(s2);
// print
console.log(s1 == '0T22KU', n1 == n, s2 == '0T22KT', n2 == n - 1);
