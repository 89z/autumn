# example 1
s1 = 'Sunday'
# example 2
s2 = "Sunday"
# example 3
s3 = '''Sunday
Monday'''
# example 4
s4 = """Sunday
Monday"""
# example 5
import io
o1 = io.StringIO('Sunday')
s5 = o1.getvalue()
# print
print([s1, s2, s3, s4, s5])
