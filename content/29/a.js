function f_date(n_date) {
   const o_date = new Date(n_date * 1000);
   const o_fmt = new Intl.DateTimeFormat('en', {
      day: 'numeric', month: 'short', weekday: 'short', year: 'numeric'
   });
   const f_parts = (m_acc, m_cur) => {
      return {...m_acc, [m_cur.type]: m_cur.value};
   };
   return o_fmt.formatToParts(o_date).reduce(f_parts, {});
}

let n = 1577858399;
let m = f_date(n);
let s = [m.weekday, m.month, m.day, m.year].join(' ');
console.log(s == 'Tue Dec 31 2019');
