String encode(Map<String, String> m) {
   var u = Uri(queryParameters: m);
   return u.query;
}

void main() {
   var m = {'west': 'left', 'east': 'right'};
   var s = encode(m);
   print(s == 'west=left&east=right');
}
