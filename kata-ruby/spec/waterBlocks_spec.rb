require_relative '../kata/waterBlocks.rb'

describe '#waterBlocks' do 
    context "Base cases" do
        it "should work" do
            input = [1, 2, 1, 4, 1, 3, 4, 3, 0, 2, 1, 3, 0, 2]
            expect(waterBlocks(input)).to eq 13
        end
    end
    context "Edge cases" do
        it "should return 0 for descending array" do
            input = [5, 4, 2, 1]
            expect(waterBlocks(input)).to eq 0
        end
        it "should return 0 for ascending array" do
            input = [0, 1, 3, 4, 6]
            expect(waterBlocks(input)).to eq 0
        end
        it "should return 0 for pyramid array" do
            input = [0, 2, 3, 5, 4, 2, 1]
            expect(waterBlocks(input)).to eq 0
        end
        it "should return 0 for flat array" do 
            input = [3, 3, 3, 3]
            expect(waterBlocks(input)).to eq 0
        end
        it "should return 0 for empty array" do
            input = []
            expect(waterBlocks(input)).to eq 0
        end
    end
end
