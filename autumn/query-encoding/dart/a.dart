String encode(Map<String, String> m) {
   var u = Uri(queryParameters: m);
   return u.query;
}

void main() {
   var m = {'month': 'March', 'day': 'Friday'};
   var s = encode(m);
   print(s == 'month=March&day=Friday');
}
