s = 'index.md'
# example 1
o1 = Time.now
File.utime(o1, o1, s)
# example 2
o2 = File.mtime(s)
# print
puts o1 == o2
