# example A
sA = 'May'
# example B
sB = "May"
# example C
sC = '''May
June'''
# example D
sD = """May
June"""
# example E
import io
o = io.StringIO('May')
sE = o.getvalue()
# print
print([sA, sB, sC, sD, sE])
