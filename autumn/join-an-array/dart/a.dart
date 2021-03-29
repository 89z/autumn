void main() {
   var a = ['West', 'East'];
   { // example 1
      var s = a.join(',');
      print(s == 'West,East');
   }
   { // example 2
      var s = a.join();
      print(s == 'WestEast');
   }
}
