require 'csv'

CSV.foreach('a.csv', headers: true) do |m|
   p m
end
