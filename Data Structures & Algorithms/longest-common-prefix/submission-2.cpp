class Solution {
public:
    string longestCommonPrefix(vector<string>& strs) {
        unordered_map<string , int>count_map;

        for(auto s : strs)
        {
            for(int i =0; i<s.size(); i++)
            {
                string put = s.substr(0 , i+1);
                count_map[put]++;
            }
        }

        int maxi =0;
        string ans = "";
        for(auto c : count_map)
        {
            cout<<c.first<<endl;
            if(maxi < c.second)
            {
                    ans = c.first;
                    maxi = c.second;
            }else if (maxi == c.second){
                if(ans.size()<c.first.size())
                {
                    ans = c.first;
                }
            }
        }
        cout<<ans<<endl;
        if(maxi == strs.size())
            return ans;


        return "";
    }
};