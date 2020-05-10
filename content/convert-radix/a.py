s_dig = '-0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ_abcdefghijklmnopqrstuvwxyz'
# example 1
def r64_encode(n_in):
   s_out = ''
   while n_in > 0:
      s_out = s_dig[n_in & 63] + s_out
      n_in >>= 6
   return s_out
# example 2
def r64_decode(s_in):
   n_out = 0
   for s_chr in s_in:
      n_out = n_out << 6 | s_dig.find(s_chr)
   return n_out
# print
import time
n1 = int(time.time())
s1 = r64_encode(n1)
n2 = r64_decode(s1)
print(n1, s1, n2 == n1)
