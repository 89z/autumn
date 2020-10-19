import argparse

parser = argparse.ArgumentParser()
parser.add_argument('--count', '-c', dest='count', type=int)
parser.add_argument('--ratio', '-r', dest='ratio', type=float)
args = parser.parse_args()

parser = argparse.ArgumentParser()
parser.add_argument('--input', '-i', dest='input', required=True)
args = parser.parse_args()

parser = argparse.ArgumentParser()
parser.add_argument('--hosts', '-H', dest='hosts', default='/etc/hosts')
args = parser.parse_args()

parser = argparse.ArgumentParser(description='Twiddles data')
parser.add_argument('--input', '-i', dest='input', help='path of input file')
args = parser.parse_args()

if not args.input:
   parser.print_help()
   sys.exit(1)
