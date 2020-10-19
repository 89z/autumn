// example 1
let o1 = {
   hours: 23,
   add: function (minutes) {
      return this.hours * 60 + minutes;
   }
};
// example 2
let o2 = {
   hours: 23,
   add(minutes) {
      return this.hours * 60 + minutes;
   }
};
// print
let n1 = o1.add(59);
let n2 = o2.add(59);
console.log(n1 == 1439, n2 == 1439);
