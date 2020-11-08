import parseopt

var
   s_start = ""
   s_end = ""
   s_in = ""
   o_opt = initOptParser()

while true:
   o_opt.next()
   if o_opt.kind == cmdEnd:
      break
   case o_opt.key
   of "s":
      s_start = o_opt.val
   of "e":
      s_end = o_opt.val
   else:
      s_in = o_opt.key

if s_in == "":
   echo """cat [flags] <input>
-s string
   start
-e string
   end"""
   quit(1)

echo s_start & s_in & s_end
