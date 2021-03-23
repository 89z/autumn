import 'dart:convert';
var src = '{"U2": {"Boy": ["Twilight", "I Will Follow"]}}';

void main() {
   var m = jsonDecode(src);
   var dst = m['U2']['Boy'][0];
   print(dst == 'Twilight');
}
