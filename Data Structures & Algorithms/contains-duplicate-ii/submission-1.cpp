class Solution {
public:
    bool containsNearbyDuplicate(vector<int>& nums, int k) {
        unordered_map<int , int>count_map;

        for(int i =0; i<nums.size(); i++)
        {
            if(count_map.find(nums[i]) != count_map.end())
            {
                if((i - count_map[nums[i]])<=k)
                    return true;
            }

            count_map[nums[i]] = i;
        }

        return false;
    }
};