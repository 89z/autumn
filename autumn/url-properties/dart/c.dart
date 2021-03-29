void main() {
   var u = Uri.parse('http://dart.dev?west=left');
   var m = u.queryParameters;
   print(m);
}
