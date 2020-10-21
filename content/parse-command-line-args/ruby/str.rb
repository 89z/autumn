require 'optparse'
s_start = ''
s_end = ''
o = OptionParser.new
o.on('-s', '--start string') { |s_val| s_start = s_val }
o.on('-e', '--end string') { |s_val| s_end = s_val }
o.parse!
if ARGV.length != 1
   o.banner = 'cat [flags] <input>'
   puts o
   exit 1
end
puts s_start + ARGV[0] + s_end
