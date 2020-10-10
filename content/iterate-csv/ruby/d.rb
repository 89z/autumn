require 'csv'

CSV.table('a.csv').each do |m|
   p m
end
