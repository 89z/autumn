>> require 'set'
>> s1 << 99 # Mutability (or s1.add(99) )
>> s1, s2 = Set[1, 2, 3, 4], [3, 4, 5, 6].to_set # different ways of creating a set
>> s1.merge(s2) # Mutability
