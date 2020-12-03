function date_f(ms_n) {
   const date_o = new Date(ms_n);
   const fmt_o = new Intl.DateTimeFormat;
   const part_f = (acc_m, cur_m) => {
      return {...acc_m, [cur_m.type]: cur_m.value};
   };
   return fmt_o.formatToParts(date_o).reduce(part_f, {});
}

let n = 1577858399;
let m = date_f(n * 1000);
let s = [m.year, m.month, m.day].join('-');
console.log(s == '2019-12-31');
