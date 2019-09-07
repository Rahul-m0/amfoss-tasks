require 'nokogiri'
=> true
irb(main):003:0> require 'open-uri'
=> true
irb(main):004:0> doc = Nokogiri::HTML(open('https://www.google.com/search?client=ubuntu&channel=fs&q=linux&ie=utf-8&oe=utf-8'))
