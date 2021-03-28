function encode(m) {
   let p = new URLSearchParams(m);
   return p.toString();
}

let m = {west: 'left', east: 'right'};
let s = encode(m);
console.log(s === 'west=left&east=right');
