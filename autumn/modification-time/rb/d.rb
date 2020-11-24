stat_o = File.stat('a.rb')
time_o = stat_o.mtime
puts time_o
