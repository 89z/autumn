let p = new URLSearchParams({month: 'March', day: 'Friday'});
let s = p.toString();
console.log(s === 'month=March&day=Friday');
