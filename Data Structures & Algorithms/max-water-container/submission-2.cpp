class Solution {
public:
    int maxArea(vector<int>& heights) {
        int left =0;
        int right = heights.size()-1;

        int maxArea =0;

        while(left < right)
        {
            int h = min(heights[left] , heights[right]);
            int w = right -left ;
            int area = h * w;
            maxArea = max(maxArea , area);

            if(heights[left]>heights[right])
            {
                right--;
            }else {
                left ++;
            }
        }

        return maxArea;

    }
};
