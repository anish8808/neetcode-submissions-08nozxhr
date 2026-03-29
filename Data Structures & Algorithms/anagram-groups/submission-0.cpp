class Solution {
public:
    vector<vector<string>> groupAnagrams(vector<string>& strs) {
        vector<vector<string>>ans;

        unordered_map<string, vector<string>>freqMap;

        for(int i =0; i<strs.size(); i++)
        {
            string str = strs[i];
            sort(str.begin() , str.end());
            freqMap[str].push_back(strs[i]);
        }

        for(auto nums : freqMap)
        {
            ans.push_back(nums.second);
        }

        return ans;
    }
};
