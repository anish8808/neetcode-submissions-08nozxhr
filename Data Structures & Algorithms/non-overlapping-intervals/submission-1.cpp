class Solution {
public:
    int eraseOverlapIntervals(vector<vector<int>>& intervals) {
        if(intervals.size()<=1)
			return 0;

		sort(intervals.begin() , intervals.end() , [](vector<int>&a , vector<int>&b){
			return a[1]<b[1];
		});


		int y = intervals[0][1];

		int count =0;

		for(int i =1; i<intervals.size(); i++)
		{
			if (y > intervals[i][0])
			{
				y = min(y  , intervals[i][1]);
				count++;
			}else {
				y = intervals[i][1];
			}
		}


		return count;
    }
};
