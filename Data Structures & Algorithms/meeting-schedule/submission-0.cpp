/**
 * Definition of Interval:
 * class Interval {
 * public:
 *     int start, end;
 *     Interval(int start, int end) {
 *         this->start = start;
 *         this->end = end;
 *     }
 * }
 */

class Solution {
public:
    bool canAttendMeetings(vector<Interval>& intervals) {
        sort(intervals.begin() , intervals.end() , [](Interval &a , Interval &b){
            return a.start <b.start;
        });

        if (intervals.size() <= 1) return true;
        int y = intervals[0].end;

        for(int i =1; i<intervals.size(); i++)
        {
            if(y>intervals[i].start)
            {
                return false;
            }else {
                y = intervals[i].end;
            }
        }

        return true;
    }
};
