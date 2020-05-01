import string
def radix64(n_in):
   s_safe = string.digits + string.ascii_letters + '-_'
   n_safe = len(s_safe)
   s_out = ''
   while n_in > 0:
      s_out = s_safe[n_in % n_safe] + s_out
      n_in //= n_safe
   return s_out
n1 = 1588306986
s1 = radix64(n1)
print(s1)
