require 'cgi'
require 'json'
require 'net/http'
res = Net::HTTP.get URI 'https://www.youtube.com/get_video_info?video_id=aqz-KE-bpKQ'
vars = CGI.parse res
json = JSON.parse vars["player_response"].first
pp json["streamingData"]["formats"]
