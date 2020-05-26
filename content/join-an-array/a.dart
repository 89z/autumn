main() {
   var s, a = ['Sun', 'Mon'];
   // example 1
   s = a.join(',');
   print(s == 'Sun,Mon');
   // example 2
   s = a.join();
   print(s == 'SunMon');
}
