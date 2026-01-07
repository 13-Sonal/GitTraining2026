def factorial(n)
	return 1 if n == 0
	result = 1
	(1..n).each do |i|
		result *= i
	end

	result
end

puts factorial(5)
