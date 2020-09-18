import sets
# example 1
var t1 = [10, 11].toHashSet
t1.incl(12)
# example 2
var t2 = ["May", "June"].toHashSet
t2.incl("July")
# print
echo t1, t2
