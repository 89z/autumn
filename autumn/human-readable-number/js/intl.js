let s = Intl.NumberFormat('en', { notation: "compact" }).format(1234);
console.log(s == '1.2K');
