import 'dart:convert';
var in_s = '{"U2": {"Boy": ["Twilight", "I Will Follow"]}}';

void main() {
   var m = jsonDecode(in_s);
   var out_s = m['U2']['Boy'][0];
   print(out_s == 'Twilight');
}
