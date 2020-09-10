o = Time.now
File.utime(o, o, 'index.md')
o2 = File.mtime('index.md')
puts o, o2 == o
