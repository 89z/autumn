class ShuffleCommand extends Command {
   final name = "shuffle";
   final description = "Shuffle and choose cards from a deck";

   ShuffleCommand() {
      argParser..addFlag('count', abbr: 'c', defaultsTo: true);
   }

   void run() {
      print(argParser.options);
   }
}

void main() {
   var o = new SuffleCommand();
}
