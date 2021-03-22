void main() {
   var a = ['May', 'June'];
   { // example 1
      var s = a.join(',');
      print(s == 'May,June');
   }
   { // example 2
      var s = a.join();
      print(s == 'MayJune');
   }
}
