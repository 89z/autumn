import string
s_dig = string.printable[:62] + '-_'
# example 1
def r64_encode(n_in):
   s_out = ''
   n_sh = n_in
   while n_sh > 0:
      n_key = n_sh & 63
      s_out = s_dig[n_key] + s_out
      n_sh >>= 6
   return s_out
# example 2
def r64_decode(s_in):
   n_out = 0
   n_key = -1
   for n_sh in range(0, 36, 6):
      s_val = s_in[n_key]
      n_out |= s_dig.find(s_val) << n_sh
      n_key -= 1
   return n_out
# print
import time
n1 = int(time.time())
s1 = r64_encode(n1)
n2 = r64_decode(s1)
print(n1, s1, n2 == n1)
