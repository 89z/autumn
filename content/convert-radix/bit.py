import string as o1
# head
def radix64(n_div):
   s_safe = o1.digits + o1.ascii_letters + '-_'
   s_out = ''
   while n_div > 0:
      s_out += s_safe[n_div & 63]
      n_div >>= 6
   return s_out
# body
s1 = radi64(1588337932)
print(s1)
