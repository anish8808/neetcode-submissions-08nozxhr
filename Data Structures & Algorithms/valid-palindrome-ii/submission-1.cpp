class Solution {
public:
    bool check (string nums , int start , int end){
        while(start<end)
        {
            if(nums[start]==nums[end])
            {
                start++;
                end --;
            }else {
                return false;
            }
        }
        return true;
        
    }
    bool validPalindrome(string nums) {
        int start =0;
        int end = nums.size()-1;
        int count =0;
        while(start<end)
        {
            if(nums[start]==nums[end])
            {
                start++;
                end --;
            }else {
                bool ans = check(nums , start+1 , end);
                bool ans1 =check(nums , start , end-1);
                return ans || ans1;
            }
        }

        return true;
    }
};