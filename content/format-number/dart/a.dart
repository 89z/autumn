import 'package:intl/intl.dart';
main() {
   var o = new NumberFormat(',###');
   var s = o.format(1000);
   print(s);
}
