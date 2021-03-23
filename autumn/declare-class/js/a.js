// example 1
let t = {
   hours: 23,
   duration: function (minutes) {
      return this.hours * 60 + minutes;
   }
};

// example 2
let u = {
   hours: 23,
   duration(minutes) {
      return this.hours * 60 + minutes;
   }
};

// print
console.log(t.duration(59) == 1439);
console.log(u.duration(59) == 1439);
