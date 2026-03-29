class Solution {
public:
    bool isValid(string s) {
        stack<char>st;

        for(int i  = 0; i <s.size(); i++)
        {
            char ch = s[i];

            //pushing into the stack

            if(ch == '(' || ch == '{' || ch == '[')
            {
                st.push(ch);
            }else {
                if(st.empty())
                {
                    return false;
                }
                //poping from the stack
                char ch2 = st.top();
                if((ch ==')' && ch2 =='(' )|| ( ch =='}' && ch2 =='{' ) || (ch==']' && ch2=='['))
                {
                    st.pop();
                }else {
                    return false;
                }

            }
            
        }

        if(st.empty())
        {
            return true;
        }

        return false;
    }
};
