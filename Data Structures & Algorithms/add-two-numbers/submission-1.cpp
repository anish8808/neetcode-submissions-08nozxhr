/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode() : val(0), next(nullptr) {}
 *     ListNode(int x) : val(x), next(nullptr) {}
 *     ListNode(int x, ListNode *next) : val(x), next(next) {}
 * };
 */

class Solution {
public:
    ListNode* addTwoNumbers(ListNode* l1, ListNode* l2) {
        int carry = 0;
        ListNode* result = new ListNode(0);
        ListNode* head = result;
        int l1val = 0;
        int l2val = 0;
        while(l1 != NULL || l2 != NULL){
            if(l1 ==NULL)
            {
                l1val =0;
            }else
            {
                l1val = l1->val;
            }

            if(l2 == NULL)
            {
                l2val =0;
            }else{
                l2val = l2->val;
            }
           
            int sum = l1val + l2val + carry ;
            carry = sum/10;
            sum = sum %10 ;
            ListNode* node1 = new ListNode(sum);
            head->next = node1;
            head = head->next;
            if(l1 != NULL)
                l1 = l1->next;
            if (l2 != NULL)
                l2 = l2->next;
        }
        
        if(carry)
        {
            ListNode* node1 = new ListNode(carry);
            head->next = node1;
        }


        return result->next ;
    
    }
};
