def fibonacci(n)
  a, b = 0, 1
  n.times do
    puts a
    a, b = b, a + b
  end
end

fibonacci(10)  # prints first 10 Fibonacci numbers

