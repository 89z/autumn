void main() {
   { // example 1
      var f = (num d, num e) {
         return d + e;
      };
      print(f(4, 5) == 9);
   }
   { // example 2
      var f = (num d, num e) => d + e;
      print(f(4, 5) == 9);
   }
}
