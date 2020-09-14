# example 1
s = 'May'
# example 2
s2 = "May"
# example 3
s3 = '''May
June'''
# example 4
s4 = """May
June"""
# example 5
import io
o = io.StringIO('May')
s5 = o.getvalue()
# print
print([s, s2, s3, s4, s5])
