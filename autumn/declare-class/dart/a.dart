class Date {
   int month = 1;
   int day = 1;
   add() {
      this.day++;
   }
}

void main() {
   var d = Date();
   d.add();
   print(d.day == 2);
}
