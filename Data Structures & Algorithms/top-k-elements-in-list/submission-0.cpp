class Solution {
public:
    vector<int> topKFrequent(vector<int>& nums, int k) {
        unordered_map<int , int>freqMap;

        for(int i =0; i<nums.size(); i++)
        {
            freqMap[nums[i]]++;
        }


        priority_queue<pair<int , int> , vector<pair<int , int>> , greater<pair<int , int>>>pq;
        for(auto i : freqMap)
        {
            pq.push({i.second , i.first});
            if(pq.size()>k)
            {
                pq.pop();
            }
        }

        vector<int>ans;

        while(!pq.empty())
        {
            auto top = pq.top();
            cout<<top.first <<top.second<<endl;
            ans.push_back(top.second);
            pq.pop();
        }


        return ans;
    }
};
