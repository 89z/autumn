import 'dart:io';
import 'package:args/args.dart';

void main(a) {
   var o = ArgParser();
   o.addOption('start', defaultsTo: '');
   o.addOption('in', defaultsTo: '');
   var m = o.parse(a);
   if (m.rest.length != 1) {
      print('match [flags] <input>\n' + o.usage);
      exit(1);
   }
   var s_start = m['start'];
   var s_in = m['in'];
   var s = m.rest[0];
   var b = s.startsWith(s_start) && s.contains(s_in);
   print(b);
}
