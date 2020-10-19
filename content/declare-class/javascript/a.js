const Radix64 = {
   digits: '-0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ_abcdefghijklmnopqrstuvwxyz',
   decode(s_in) {
      let n_out = 0;
      for (let s_chr of s_in) {
         n_out = n_out * 64 + this.digits.indexOf(s_chr);
      }
      return n_out;
   }
};

const n = Radix64.decode('0UYVwl');
console.log(n);
