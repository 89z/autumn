let in_s = '{"U2": {"Boy": ["Twilight", "I Will Follow"]}}';
let m = JSON.parse(in_s);
let out_s = m.U2.Boy[0];
console.log(out_s == 'Twilight');
