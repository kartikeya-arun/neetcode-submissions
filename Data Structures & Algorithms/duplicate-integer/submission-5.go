func hasDuplicate(nums []int) bool {
    seen := make(map[int]int)
    for i:=0;i<len(nums);i++{
        _,ok := seen[nums[i]]
        if ok{
            return true
        }else{
            seen[nums[i]] = 1
        }
    }
    return false
}
