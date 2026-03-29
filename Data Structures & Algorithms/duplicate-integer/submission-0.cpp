class Solution {
public:
    bool hasDuplicate(vector<int>& nums) {
        unordered_map<int , bool > freqMap;

        for(int i =0; i<nums.size(); i++)
        {
            if(freqMap.find(nums[i])!=freqMap.end())
            {
                return true;
            }
            freqMap[nums[i]] = true;
        }

        return false;
    }
};