require 'getoptlong'
opts = GetoptLong.new(
   ['--start', '-s', GetoptLong::REQUIRED_ARGUMENT],
   ['--end', '-e', GetoptLong::REQUIRED_ARGUMENT]
)
opts.each { |option, value|
   case option
   when '--start'
      s_start = value
   when '--end'
      s_end = value
   end
}

if ARGV.length
end
