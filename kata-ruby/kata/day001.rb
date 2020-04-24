# Write a function that takes in a string of one or more words, and returns the same string, but with all five or more letter words reversed (Just like the name of this Kata). Strings passed in will consist of only letters and spaces. Spaces will be included only when more than one word is present.

def reverse(str)
  acc = ""
  str.length().times do |i| # i = 0 -> i = length - 1
    acc += str.slice(str.length() - i - 1)
  end
  return acc
end

def spinWords(string)
  words = string.split(' ')
  words.each_with_index do |word, i|
    if word.length() >= 5
      words[i] = reverse(word)
    end
  end
  return words.join(' ')
end