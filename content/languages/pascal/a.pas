uses fphttpclient;
var m1: TFPHTTPClient;
var s1: string;
begin
   m1 := TFPHTTPClient.Create(nil);
   s1 := m1.get('https://speedtest.lax.hivelocity.net');
   writeln(s1);
end.
