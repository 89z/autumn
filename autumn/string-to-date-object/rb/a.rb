require 'date'
s = "2011-05-03 10:00:00"
fmt = "%Y-%m-%d %H:%M:%S"
t = DateTime.strptime(s, fmt).to_time
