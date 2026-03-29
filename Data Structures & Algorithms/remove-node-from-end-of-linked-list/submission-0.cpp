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
    int GetLength(ListNode* head)
    {
        if(head==NULL)
            return 0;

        ListNode* temp = head;
        int lenght = 0;
        while(temp!=NULL)
        {
            lenght++;
            temp = temp->next;
        }

        return lenght;
    }


    ListNode* removeNthFromEnd(ListNode* head, int n) {
        if(head==NULL)
            return head;


        int len = GetLength(head);

        if(len==n)
        {
            head = head->next;
            return head;
        }

        int removelen = len - n-1;

        ListNode* temp = head;

        while(removelen>0)
        {
            temp = temp->next;
            removelen--;
        }


        if(temp->next == NULL)
        {
            delete(temp);
        }
        else {
            ListNode* del = temp->next;
            temp->next = temp->next->next;
            delete(del);
        }
        return head;
    }
};
