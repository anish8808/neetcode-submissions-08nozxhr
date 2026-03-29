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
    ListNode* getMiddleNode(ListNode* head)
    {
        if(head==NULL)
            return NULL;

        ListNode* fast = head;
        ListNode* slow = head;

        while(fast!=NULL && fast->next!=NULL && fast->next->next !=NULL)
        {
            slow = slow->next;
            fast = fast->next->next;
        }

        return slow;
    }

    ListNode* reverseNode(ListNode* head)
    {
        if(head==NULL)
            return NULL;

        ListNode* curr = head;
        ListNode* prev = NULL;
        ListNode* next = NULL;

        while(curr != NULL)
        {
            next = curr->next;
            curr->next = prev ;
            prev = curr;
            curr = next;
        }

        return prev;
    }

    void reorderList(ListNode* head) {
      ListNode* middle = getMiddleNode(head);  
      ListNode* start1 = reverseNode(middle->next);
      middle->next = NULL;
      ListNode* start = head;
      while(start1 !=NULL )
      {
            ListNode* next = start->next;
            ListNode* next1 = start1->next;
            start->next = start1;
            start1->next = next;
            start = next;
            start1 = next1;
      }
    }
};
