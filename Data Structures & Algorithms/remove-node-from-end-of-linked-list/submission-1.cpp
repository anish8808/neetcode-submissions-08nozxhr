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
    ListNode* removeNthFromEnd(ListNode* head, int n) {
        if(n==0 || head ==NULL)
            return head;

        ListNode* dummy = new ListNode(0);
        dummy->next = head;
        ListNode* startNode = dummy;
        while(n>0)
        {
            startNode = startNode->next;
            n--;
        }

        ListNode* prev = dummy;
        while(startNode->next!=NULL && startNode !=NULL)
        {
            prev = prev->next;
            startNode = startNode->next;
        }


        if(prev->next->next==NULL)
        {
            prev->next = NULL;
        }
        else {
            prev->next = prev->next->next;
        }


        return dummy->next;
    }
};
