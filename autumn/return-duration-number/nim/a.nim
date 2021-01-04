import times

proc sinceHours(s: string): int64 =
   let t = parse(s, "yyyy-MM-dd")
   let u = now()
   return (u - t).inHours

let n = sinceHours("2020-12-31")
echo n
