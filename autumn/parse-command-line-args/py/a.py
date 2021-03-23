import argparse

arg = argparse.ArgumentParser()
arg.add_argument('-p', default='', help='prefix')
arg.add_argument('-s', default='', help='suffix')
arg.add_argument('stem')
parse = arg.parse_args()

print(parse.p + parse.stem + parse.s)
