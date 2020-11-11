import argparse
in_o = argparse.ArgumentParser()
in_o.add_argument('-p', default='', help='prefix')
in_o.add_argument('-s', default='', help='suffix')
in_o.add_argument('stem')
out_o = in_o.parse_args()
print(out_o.p + out_o.stem + out_o.s)
