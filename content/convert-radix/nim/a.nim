type Radix36 = object
   digit: string

proc newRadix36(): Radix36 =
   return Radix36(digit: "0123456789abcdefghijklmnopqrstuvwxyz")

proc encode(o: Radix36, n: int): string =
   var
      s = ""
      t = n
   while true:
      s = o.digit[t mod 36] & s
      t = t div 36
      if t == 0: break
   return s

proc decode(o: Radix36, s: string): int =
   var n: int
   for c in s: n = n * 36 + o.digit.find(c)
   return n

let o = newRadix36()
let s = o.encode(1577858399)
let n = o.decode("q3ezbz")
echo s == "q3ezbz" and n == 1577858399
