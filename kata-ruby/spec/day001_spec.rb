require_relative '../kata/day001.rb'

describe '#spinWords' do
  context "Base cases" do
    it "should spin only 5 letter words" do
      expect(spinWords("it is wednesday my dudes")).to eq "it is yadsendew my sedud"
    end
    it "should maintain capitalization" do
      expect(spinWords("It is Wednesday my dudes")).to eq "It is yadsendeW my sedud"
    end
    it "should work with only short words" do
      expect(spinWords("i am a cat")).to eq "i am a cat"
    end
    it "should work with only long words" do
      expect(spinWords("helicopter airplane balloon")).to eq "retpocileh enalpria noollab"
    end
  end
end