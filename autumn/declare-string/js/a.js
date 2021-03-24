let s;

// slash
s = String.raw`south\north`;
console.log(s);

// quote
s = String.raw`south"north`;
console.log(s);
s = String.raw`south'north`;
console.log(s);

// newline
s = String.raw`south
north`;
console.log(s);
