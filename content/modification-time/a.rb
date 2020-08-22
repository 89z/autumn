o1 = Time.now
File.utime(o1, o1, 'index.md')
o2 = File.mtime('index.md')
puts o1, o2 == o1
