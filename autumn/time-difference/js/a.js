function date_f(ms_n) {
   const date_o = new Date(ms_n);
   const fmt_o = new Intl.DateTimeFormat('en', {
      hourCycle: 'h23', timeStyle: 'full', timeZone: 'UTC'
   });
   const part_f = (acc_m, cur_m) => {
      return {...acc_m, [cur_m.type]: cur_m.value};
   };
   return fmt_o.formatToParts(date_o).reduce(part_f, {});
}

let n = Date.parse('2019-12-31T00:00:00');
let n2 = Date.parse('2019-12-31T23:59:59');
let m = date_f(n2 - n);
let s = [m.hour, m.minute, m.second].join(':');
console.log(s == '23:59:59');
