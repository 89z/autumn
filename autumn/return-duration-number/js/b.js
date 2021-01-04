function since(s) {
   let n = Date.now() - Date.parse(s);
   return n / 1000;
}

let n = since('2020-12-31');
console.log(n);
