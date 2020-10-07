require 'csv'

CSV.foreach("foo.csv") do |row|
  puts row.join("\t")
end
