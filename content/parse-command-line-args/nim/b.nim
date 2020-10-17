import parseopt, strutils

var
   n_len = 1
   n_start = 0
   o_opt = initOptParser()
   s_in = ""

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

if s_in == "":
   echo """slice [flags] <string>
--start int
   starting position
--len int
   substring length (default 1)"""
   quit(1)

echo s_in[n_start ..< n_start + n_len]
