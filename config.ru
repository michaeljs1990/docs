#!/usr/bin/env ruby
require 'rubygems'
require 'gollum/app'

gollum_path = File.expand_path(File.dirname(__FILE__))

Precious::App.set(:gollum_path, gollum_path)
Precious::App.set(:wiki_options, index_page: "README")

run Precious::App
