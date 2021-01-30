void main() {
   var u = Uri(queryParameters: {'month': 'March', 'day': 'Friday'});
   print(u.query == 'month=March&day=Friday');
}
