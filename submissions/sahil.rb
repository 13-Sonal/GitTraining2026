def reverse_string(string)
    reversed_string = ""
    
    string.each_char do |char|
        reversed_string = char + reversed_string
    end 
    reversed_string
end

def palindrome?(string)
    string == string.reverse 
end

def fibonacci(n) 
    a, b = 0, 1
    puts a
    (n-1).times do 
        a, b = b, a + b
        puts a
    end
end