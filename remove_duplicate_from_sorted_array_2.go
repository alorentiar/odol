func removeDuplicates(nums []int) int {
    myMap := make(map[int]int)
    count := 0

    for i := 0; i < len(nums); i++{
        if myMap[nums[i]] < 2 {
            myMap[nums[i]] = myMap[nums[i]] + 1
            nums[count] = nums[i]
            count++
        }
    }

    return count
}