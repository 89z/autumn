function date_f(ms_n) {
   const date_o = new Date(ms_n);
   const fmt_o = new Intl.DateTimeFormat('en', {
      day: 'numeric', month: 'short', weekday: 'short', year: 'numeric'
   });
   const part_f = (acc_m, cur_m) => {
      return {...acc_m, [cur_m.type]: cur_m.value};
   };
   return fmt_o.formatToParts(date_o).reduce(part_f, {});
}

let n = 1577858399;
let m = date_f(n * 1000);
let s = [m.weekday, m.month, m.day, m.year].join(' ');
console.log(s == 'Tue Dec 31 2019');
