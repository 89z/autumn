import tempfile
# create and remove
m1 = tempfile.NamedTemporaryFile()
s1 = m1.name
# create only
m2 = tempfile.NamedTemporaryFile(delete = 0)
s2 = m2.name
# print
print(s1, s2)
