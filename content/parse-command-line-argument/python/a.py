import argparse

parser = argparse.ArgumentParser()
parser.add_argument('--count', '-c',
  dest='count', type=int)
parser.add_argument('--ratio', '-r',
  dest='ratio', type=float)
args = parser.parse_args()

if args.count:
  print('Count: %d' % args.count)
if args.ratio:
  print('Ratio: %.3f' % args.ratio)

parser = argparse.ArgumentParser()
parser.add_argument('--input', '-i',
  dest='input', required=True)
args = parser.parse_args()

f = open(args.input)

parser = argparse.ArgumentParser()
parser.add_argument('--hosts', '-H',
  dest='hosts', default='/etc/hosts')
args = parser.parse_args()

f = open(args.hosts)

parser = argparse.ArgumentParser(
  description='Twiddles data')
parser.add_argument('--input', '-i',
  dest='input', help='path of input file')
args = parser.parse_args()

if not args.input:
  parser.print_help()
  sys.exit(1)

from optparse import OptionParser
parser = OptionParser()
parser.add_option("-f", "--file", dest="filename",
                  help="write report to FILE", metavar="FILE")
parser.add_option("-q", "--quiet",
                  action="store_false", dest="verbose", default=True,
                  help="don't print status messages to stdout")
 
(options, args) = parser.parse_args()
