function Radix64() {
   let s = '-0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ_abcdefghijklmnopqrstuvwxyz';
   this.s_dig = s;
}

Radix64.prototype.encode = function(n_in) {
   let s_out = '';
   do {
      s_out = this.s_dig[n_in % 64] + s_out;
      n_in = Math.trunc(n_in / 64);
   } while (n_in > 0);
   return s_out;
};

Radix64.prototype.decode = function(s_in) {
   let n_out = 0;
   for (let s_chr of s_in) {
      n_out = n_out * 64 + this.s_dig.indexOf(s_chr);
   }
   return n_out;
};

let n = Math.trunc(new Date / 1000);
let o = new Radix64;
let s = o.encode(n);
let n2 = o.decode(s);
console.log(n, s, n2 == n);
