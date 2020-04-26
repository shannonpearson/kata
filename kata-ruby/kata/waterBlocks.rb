def waterBlocks(arr)
    rightMax = 0;
    leftMax = 0;
    left = []
    right = []

    arr.length().times do |i|
        left << leftMax
        right.unshift(rightMax)
        if (arr[i] > leftMax)
            leftMax = arr[i]
        end
        if (arr[arr.length() - i - 1] > rightMax)
            rightMax = arr[arr.length() - i - 1]
        end
    end

    result = []
    arr.length().times do |i|
        result << ([left[i], right[i]].min - arr[i] > 0 ? [left[i], right[i]].min - arr[i] : 0)
    end
    return result.sum
end

waterBlocks([1, 2, 1, 4, 1, 3, 4, 3, 0, 2, 1, 3, 0, 2])