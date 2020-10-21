require 'getoptlong'

s_start = ''
s_end = ''

GetoptLong.new(
   ['--start', '-s', GetoptLong::REQUIRED_ARGUMENT],
   ['--end', '-e', GetoptLong::REQUIRED_ARGUMENT]
).each { |s_key, s_val|
   case s_key
   when '--start'
      s_start = s_val
   when '--end'
      s_end = s_val
   end
}

if ARGV.length != 1
   puts <<eof
cat [flags] <input>
-s string
   start
-e string
   end
eof
   exit 1
end

puts s_start + ARGV[0] + s_end
