function format(date_o) {
   function part(acc_m, cur_m) {
      return { ...acc_m, [cur_m.type]: cur_m.value };
   }
   let fmt_o = new Intl.DateTimeFormat('en', {
      year: 'numeric', month: '2-digit', day: '2-digit'
   });
   let date_m = fmt_o.formatToParts(date_o).reduce(part, {});
   return [date_m.year, date_m.month, date_m.day].join('-');
}

let s = format(new Date);
console.log(s);
