void main() {
   var n = 1577858399;
   var o = new DateTime.fromMillisecondsSinceEpoch(n * 1000);
   // example 1
   var s1 = o.toString();
   // example 2
   var s2 = [o.year, o.month, o.day, o.weekday].join('-');
   // print
   print(s1 == '2019-12-31 23:59:59.000' && s2 == '2019-12-31-2');
}
