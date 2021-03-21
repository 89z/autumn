void main() {
   { // example 1
      Map<String, int> m;
      print(m);
   }
   { // example 2
      var m = <String, int>{};
      print(m);
   }
   { // example 3
      var m = Map<String, int>();
      print(m);
   }
   { // example 4
      var m = {'month': 12, 'day': 31};
      print(m);
   }
}
