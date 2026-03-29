func max(a , b int) int{
    if a>=b{
        return a
    } 
    return b
}

func merge(intervals [][]int) [][]int {
    var ans [][]int 

    if len(intervals)<=1 {
        return intervals
    }

    sort.Slice(intervals , func(i , j int)bool {
        return intervals[i][0]<intervals[j][0]
    })

    x , y := intervals[0][0] , intervals[0][1]
    for i:=1; i<len(intervals); i++{
        if y>=intervals[i][0]{
            y = max(y  , intervals[i][1])
            continue
        }
        ans = append(ans , []int{x , y})
        x = intervals[i][0]
        y = intervals[i][1]
    }

    ans = append(ans , []int{x , y})

    return ans 
}
