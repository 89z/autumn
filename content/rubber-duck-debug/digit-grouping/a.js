function f_group(s_in) {
   let n_neg = ('-' + s_in).lastIndexOf('-');
   let n_dot = (s_in + '.').indexOf('.');
   let s_out = '';
   for (let [n_pos, s_pos] of Object.entries(s_in)) {
      if (n_neg < n_pos && n_pos < n_dot && n_pos % 3 == n_dot % 3) {
         s_out += ',';
      }
      s_out += s_pos;
   }
   return s_out;
}
let a1 = ['123', '1234', '-123', '-1234', '123.4', '1234.5'];
for (let s1 of a1) {
   console.log(f_group(s1));
}
