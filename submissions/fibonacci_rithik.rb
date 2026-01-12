n = 10 
a = 0
b = 1
n.times do
  puts a
  a, b = b, a + b
end
