function f() {
   console.clear();
   let new_n = Date.now();
   let sec_n = (new_n - old_n) / 1000;
   let s = sec_n.toFixed(2);
   console.log(s);
}

let old_n = Date.now();
setInterval(f, 10);
