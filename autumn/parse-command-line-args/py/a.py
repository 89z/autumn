import argparse
o_in = argparse.ArgumentParser()
o_in.add_argument('-p', default='', help='prefix')
o_in.add_argument('-s', default='', help='suffix')
o_in.add_argument('stem')
o_out = o_in.parse_args()
print(o_out.p + o_out.stem + o_out.s)
