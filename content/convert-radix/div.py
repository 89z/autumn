import string
# head
def radix64(n_div):
   s_safe = string.digits + string.ascii_letters + '-_'
   s_out = ''
   while n_div > 0:
      n_div, n_mod = divmod(n_div, 64)
      s_out += s_safe[n_mod]
   return s_out
# body
s1 = radix64(1588337932)
print(s1)
