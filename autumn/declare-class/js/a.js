let d;

// example 1
d = {
   month: 12,
   day: 31,
   year: function () {
      this.month = 1;
      this.day = 1;
   }
};

d.year();
console.log(d.day == 1);

// example 2
d = {
   month: 12,
   day: 31,
   year() {
      this.month = 1;
      this.day = 1;
   }
};

d.year();
console.log(d.day == 1);
