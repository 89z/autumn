void main() {
   var u = Uri.parse('http://dart.dev?month=May&day=Friday');
   var m = u.queryParameters;
   print(m);
}
