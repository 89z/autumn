let d;

// example 1
d = {
   month: 1,
   day: 1,
   add: function () {
      this.day++;
   }
};

d.add();
console.log(d.day == 2);

// example 2
d = {
   month: 1,
   day: 1,
   add() {
      this.day++;
   }
};

d.add();
console.log(d.day == 2);
