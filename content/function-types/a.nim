proc f(s: string): int =
   return s.len

let n = f("May")
echo n == 3
