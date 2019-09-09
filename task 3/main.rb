require 'nokogiri'
require 'open-uri'
result = Nokogiri::HTML(open('https://www.google.com/search?q=linux'))
result.css('div').each do |link|
puts link.text
puts ""
end
