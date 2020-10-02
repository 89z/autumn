require 'uri'

url = "http://google.com:80/foo?q=3#bar"
up = URI(url)

protocol = up.scheme
hostname = up.host
port = up.port
path = up.path
query_str = up.query
fragment = up.fragment

# Ruby 1.9; returns array of pairs:
params = URI.decode_www_form(query_str)
