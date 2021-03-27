class Date {
   int month = 12;
   int day = 31;
   year() {
      this.month = 1;
      this.day = 1;
   }
}

void main() {
   var d = Date();
   d.year();
   print(d.day == 1);
}
