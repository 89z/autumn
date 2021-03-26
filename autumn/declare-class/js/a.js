let t, n;

// example 1
t = {
   hours: 23,
   duration: function (minutes) {
      return this.hours * 60 + minutes;
   }
};

n = t.duration(59);
console.log(n == 1439);

// example 2
t = {
   hours: 23,
   duration(minutes) {
      return this.hours * 60 + minutes;
   }
};

n = t.duration(59);
console.log(n == 1439);
