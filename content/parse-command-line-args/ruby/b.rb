require 'optparse'
m = {:start => '', :end => ''}
o = OptionParser.new
o.on('--start string')
o.on('--end string')
o.parse!(into: m)

if ARGV.length != 1
   o.banner = 'cat [flags] <input>'
   puts o
   exit 1
end

puts m[:start] + ARGV[0] + m[:end]
