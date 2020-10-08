require 'csv'

CSV.foreach('a.csv').with_index do |a_row, n_row|
   print n_row, a_row, "\n"
end
