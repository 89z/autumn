require 'optparse'
options = {}
OptionParser.new do |opts|
   opts.on('-i', '--input PATH') do |arg|
      options[:input] = arg
   end
end.parse!

if options[:input].nil?
   f = $stdin
else
   f = File.open(options[:input])
end
OptionParser.new do |opts|
  opts.banner =
    "usage: #{$PROGRAM_NAME} [OPTIONS] [ARG ...]"
  opts.on('-i', '--input PATH') do |arg|
    options[:input] = arg
  end
  if options[:input].nil?
    puts opts
    exit 1
  end
end.parse!

require 'optparse'
sflag = false
longflag = false
count = 0
percent = 50
fruit = nil
OptionParser.new do |opts|
  opts.banner += " [arguments...]"
  opts.separator "This demo prints the results of parsing the options."
  opts.version = "0.0.1"
  opts.on("-s", "Enable short flag") {sflag = true}
  opts.on("--long", "Enable long flag") {longflag = true}
  opts.on("-b", "--both", "Enable both -s and --long") {sflag = true; longflag = true}
  opts.on("-c", "--count", "Add 1 to count") {count += 1}
  opts.on("-p", "--percent PERCENT", /[0-9]+%?/i, "Percent [50%]") {|arg| percent = arg.to_i}
  opts.on("-f", "--fruit FRUIT", [:apple, :banana, :orange, :pear], "Fruit (apple, banana, orange, pear)") {|arg| fruit = arg}
  begin
    opts.parse!
  rescue OptionParser::ParseError => error
    $stderr.puts error
    $stderr.puts "(-h or --help will show valid options)"
    exit 1
  end
end
 
print <<EOF
Short flag: #{sflag}
Long flag: #{longflag}
Count: #{count}
Percent: #{percent}%
Fruit: #{fruit}
Arguments: #{ARGV.inspect}
EOF
