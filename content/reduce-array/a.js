let a = ['May', 'June'];
let f = (s_acc, s_cur) => s_acc + s_cur;
// example 1
let s = a.reduce(f);
// example 2
let s2 = '';
for (let s_cur of a) {
   s2 = f(s2, s_cur);
}
// print
console.log(s == 'MayJune', s2 == 'MayJune');
