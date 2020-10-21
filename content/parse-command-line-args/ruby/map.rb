require 'optparse'
o = OptionParser.new do |opts|
   opts.banner = 'cat [flags] <input>'
   opts.on('-s', '--start string')
   opts.on('-e', '--end string')
end
m = {:start => '', :end => ''}
o.parse!(into: m)
if ARGV.length != 1
   puts o
   exit 1
end
puts m[:start] + ARGV[0] + m[:end]
