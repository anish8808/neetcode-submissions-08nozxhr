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

    ListNode* reverseList(ListNode* head)
    {
        if (head == NULL)
            return NULL;

        ListNode* curr = head;
        ListNode* prev = NULL;
        ListNode* next = NULL;

        while(curr)
        {
            next = curr->next;
            curr->next = prev;
            prev = curr;
            curr = next;
        }

        return prev;
    }

    ListNode* reverseBetween(ListNode* head, int left, int right) {

        if(head == NULL)
            return NULL;

        if(left == right)
        {
            return head;
        }

        ListNode * dummy = new ListNode(0);
        dummy->next = head;
        ListNode* prev = dummy;
        for(int i =1; i<left; i++)
            prev = prev->next;

        ListNode* leftNode = prev->next;
        ListNode* rightNode = leftNode;

        for(int i = left; i<right; i++)
        {
            rightNode = rightNode->next;
        }

       
        
        ListNode* secondList = rightNode->next;
        rightNode->next = NULL;
        
        ListNode* startNode = reverseList(leftNode);

        prev->next= startNode;
        leftNode->next = secondList;
        return dummy->next;
    }

    
};