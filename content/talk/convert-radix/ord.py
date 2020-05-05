s_enc = '0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ-_'
s_dec = '_~~!"#$%&\'()*~~~~~~~EFGHIJKLMNOPQRSTUVWXYZ[\]^~~~~`~+,-./0123456789:;<=>?@ABCD'

def r64_encode(n_in):
   s_out = ''
   while n_in > 0:
      s_out = s_enc[n_in & 63] + s_out
      n_in >>= 6
   return s_out

def r64_decode(s_in):
   n_out = 0
   for s_chr in s_in:
      n_ord = ord(s_chr) - 45
      s_out = s_dec[n_ord]
      n_out = n_out << 6 | ord(s_out) - 33
   return n_out

import time
n1 = int(time.time())
s1 = r64_encode(n1)
n2 = r64_decode(s1)
print(n1, s1, n2 == n1)
