import 'package:args/args.dart';

void main(a) {
   var o = ArgParser();
   o.addOption('start');
   o.addOption('len', defaultsTo: 1);
   var m = o.parse(a);
   print(m);
}
