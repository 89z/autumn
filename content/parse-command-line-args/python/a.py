import argparse
o_in = argparse.ArgumentParser()
o_in.add_argument('--start', default='')
o_in.add_argument('--end', default='')
o_in.add_argument('input')
o_out = o_in.parse_args()
print(o_out.start + o_out.input + o_out.end)
