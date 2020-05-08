import os
s1 = os.path.abspath('index.md')
s2 = os.getenv('HOMEDRIVE')
# example 1
s3 = os.path.join(s2, 'a.md')
os.symlink(s1, s3)
# example 2
os.chdir(s1)
os.symlink(s1, 'b.md')
