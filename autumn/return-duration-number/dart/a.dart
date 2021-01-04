num sinceHours(String s) {
   var t = DateTime.parse(s);
   return DateTime.now().difference(t).inHours;
}

void main() {
   var n = sinceHours('2020-12-31');
   print(n);
}
