import re
s_in = 'Wednesday'
s_re = 'e.'
# example 1
o = re.search(s_re, s_in)
s = o.group()
# example 2
a = re.findall(s_re, s_in)
# print
print(s, a)
