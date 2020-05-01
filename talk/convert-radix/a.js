function f1(ni, nr) {
   let s1 = '0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ';
   let s2 = '';
   while (ni > 0) {
      s2 = s1[ni % nr] + s2;
      ni = Math.floor(ni / nr);
   }
   return s2;
}
