function f_date(n_ms) {
   const o_date = new Date(n_ms);
   const o_fmt = new Intl.DateTimeFormat('en', {
      hourCycle: 'h23', timeStyle: 'full', timeZone: 'UTC'
   });
   const f_parts = (m_acc, m_cur) => {
      return {...m_acc, [m_cur.type]: m_cur.value};
   };
   return o_fmt.formatToParts(o_date).reduce(f_parts, {});
}

let n = Date.parse('2019-12-31T00:00:00');
let n1 = Date.parse('2019-12-31T23:59:59');
let m = f_date(n1 - n);
let s = [m.hour, m.minute, m.second].join(':');
console.log(s == '23:59:59');
