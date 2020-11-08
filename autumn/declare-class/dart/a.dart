class Time {
   final hours = 23;
   num duration(num minutes) {
      return hours * 60 + minutes;
   }
}

void main() {
   var o = new Time();
   var n = o.duration(59);
   print(n == 1439);
}
