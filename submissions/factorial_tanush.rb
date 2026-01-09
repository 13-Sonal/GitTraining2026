def factorial(n)
  return 1 if n <= 1
  (1..n).inject(:*)  # multiply all numbers from 1 to n
end

puts factorial(5)  # 120
puts factorial(0)  # 1

