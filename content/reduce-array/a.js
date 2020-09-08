let a = ['May', 'June'];

// example 1
let f = (s_acc, s_cur) => s_acc + s_cur;
let s = a.reduce(f);

// example 2
let f2 = function (m_acc, s_cur, n_idx) {
   m_acc[s_cur] = n_idx;
   return m_acc;
};
let m = a.reduce(f2, {});

// print
console.log(s, m);
