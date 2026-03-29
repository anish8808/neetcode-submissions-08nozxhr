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
    ListNode* mergeTwoSortedList(ListNode* list1 , ListNode* list2)
    {
        if(list1 == NULL)
            return list2;
        if(list2 == NULL)
            return list1;

        ListNode* dummyNode = new ListNode(0);
        ListNode* curr = dummyNode;
        while(list1 != NULL && list2 != NULL)
        {
            if(list1->val >= list2->val)
            {
                curr->next = list2;
                list2 = list2->next;
            }else {
                curr->next = list1;
                list1 = list1->next;
            }

            curr = curr->next;
        }

        if(list1 != NULL)
        {
            curr->next = list1;
        }

        if(list2 != NULL)
        {
            curr->next = list2;
        }
        ListNode *result = dummyNode->next;
        delete dummyNode;
        return result ;
        
    }
    ListNode* mergeKLists(vector<ListNode*>& lists) {
        ListNode* list1 = NULL;
        for(int i =0; i<lists.size(); i++)
        {
            list1 = mergeTwoSortedList(list1 , lists[i]);
        }

        return list1;
    }
};
