import 'dart:io' show Platform;
main() {
   Map<String, String> envVars = Platform.environment;
   print(envVars['BROWSER']);
}
