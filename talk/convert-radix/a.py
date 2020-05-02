import string

def radix64(n_div):
   s_safe = string.digits + string.ascii_letters + '-_'
   n_safe = len(s_safe)
   s_out = ''
   while n_div > 0:
      n_div, n_mod = divmod(n_div, n_safe)
      s_out = s_safe[n_mod] + s_out
   return s_out

n1 = 1588337932
s1 = radix64(n1)
print(s1)
