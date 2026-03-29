class Solution {
public:
    int lengthOfLongestSubstring(string s) {
        unordered_map<int , int> freqMap;
        int left = 0;
        int longest = 0;
        for(int right =0; right<s.size(); right++)
        {
            while(freqMap.find(s[right]) != freqMap.end())
            {
                freqMap.erase(s[left]);
                left++;
            }

            freqMap[s[right]]++;

            longest = max(longest , right -left +1);

        }


        return longest;
    }
};
