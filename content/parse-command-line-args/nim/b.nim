import parseopt, strutils

var
   n_len = 1
   n_start: int
   o_opt = initOptParser()
   s_in: string

while true:
   o_opt.next()
   if o_opt.kind == cmdEnd:
      break
   case o_opt.key
   of "start":
      n_start = o_opt.val.parseInt
   of "len":
      n_len = o_opt.val.parseInt
   else:
      s_in = o_opt.key

let s_out = s_in[n_start ..< n_start + n_len]
echo s_out
