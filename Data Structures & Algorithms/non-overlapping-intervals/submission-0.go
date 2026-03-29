func min(x , y int) int {
	if x<=y{
		return x
	}

	return y
}

func eraseOverlapIntervals(intervals [][]int) int {
    sort.Slice(intervals , func (i , j int) bool {
		return intervals[i][1]<intervals[j][1]
	})

	if len(intervals)<=1{
		return 0
	}

	var count int
	y := intervals[0][1]

	for i:=1; i<len(intervals); i++{
		if(y>intervals[i][0]){

			y = min(y , intervals[i][1])
			count++
		}else {
			y = intervals[i][1]
		}
	}

	return count
}
