function join(t, a, s) {
   function format(m) {
      let f = new Intl.DateTimeFormat('en', m);
      return f.format(t);
   }
   return a.map(format).join(s);
}

let a = [
   {weekday: 'short'}, {month: 'short'}, {day: 'numeric'}, {year: 'numeric'}
];

let s = join(new Date, a, ' ');
console.log(s);
