class Solution {
public:
    int characterReplacement(string s, int k) {
        unordered_map<int , int>freqMap;

        int maxfreq = 0;
        int longest = 0;
        int left = 0;

        for(int right =0; right<s.size(); right++)
        {
            freqMap[s[right]]++;

            maxfreq = max(maxfreq , freqMap[s[right]]);

            while((right -left +1 ) - maxfreq >k)
            {
                freqMap[s[left]]--;
                left++;
            }

            longest = max(longest , right-left+1);
            
        }

        return longest;
    }
};
