import 'dart:io';
import 'package:args/args.dart';

void main(a) {
   var o = ArgParser();
   o.addOption('start', defaultsTo: '0');
   o.addOption('len', defaultsTo: '1');
   var m = o.parse(a);

   if (m.rest.length != 1) {
      print('slice [flags] <input>\n' + o.usage);
      exit(1);
   }

   var n_start = int.parse(m['start']);
   var n_len = int.parse(m['len']);
   var s_out = m.rest[0].substring(n_start, n_start + n_len);
   print(s_out);
}
