void main() {
   var n = 7654321;
   var s = n.toString();
   var s2;
   var n2 = s.length();
   for (auto s2a: s) {
      if (n2 % 3 == 0) {
         s2 += ",";
      }
      s2 += s2a;
      n2--;
   }
}
