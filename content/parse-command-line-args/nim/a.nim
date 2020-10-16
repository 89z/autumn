import os, parseopt
var argCtr: int
for kind, key, value in getOpt():
   case kind
   of cmdArgument:
      echo "Got arg ", argCtr, ": \"", key, "\""
      argCtr.inc
   of cmdLongOption, cmdShortOption:
      case key
      of "v", "n", "z", "w":
         echo "Got a \"", key, "\" option with value: \"", value, "\""
      else:
      echo "Unknown option: ", key
   of cmdEnd:
      discard
