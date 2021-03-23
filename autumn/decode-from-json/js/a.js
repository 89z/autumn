let src = '{"U2": {"Boy": ["Twilight", "I Will Follow"]}}';
let m = JSON.parse(src);
let dst = m.U2.Boy[0];
console.log(dst == 'Twilight');
