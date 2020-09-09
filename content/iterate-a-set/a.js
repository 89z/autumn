let t = new Set(['May', 'June']);
// example 1
for (let s of t) {
   console.log(s);
}
// example 2
let f = s => console.log(s);
t.forEach(f);
