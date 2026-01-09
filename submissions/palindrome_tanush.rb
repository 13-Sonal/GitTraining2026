def palindrome?(str)
  str = str.downcase.gsub(/\W/, '')  # remove non-alphanumeric chars
  str == str.reverse
end

puts palindrome?("Racecar")   # true
puts palindrome?("Hello")     # false

