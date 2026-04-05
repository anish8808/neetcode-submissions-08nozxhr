class Solution {
public:
    int calPoints(vector<string>& operations) {
        int totalSum =0;
        stack<int>st;
        for (int i =0; i<operations.size(); i++)
        {
            if (operations[i]=="+")
            {
                int top1 = st.top();
                st.pop();
                int top2 = st.top();
                st.push(top1);
                st.push(top1+top2);
            }else if (operations[i]=="C"){
                st.pop();
            }else if (operations[i]=="D"){
                st.push(2 * st.top());
            }else {
                int a = stoi(operations[i]);
                st.push(a);
            }
        }

        while(!st.empty())
        {
            int top = st.top();
            totalSum +=top;
            st.pop();
        }

        return totalSum;
    }
};