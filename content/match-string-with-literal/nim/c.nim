proc startsWith(s: string, c: char): bool =
   return s[0] == c

let s = "June"
echo startsWith(s, 'J')
