/*
// Definition for a Node.
class Node {
public:
    int val;
    Node* next;
    Node* random;
    
    Node(int _val) {
        val = _val;
        next = NULL;
        random = NULL;
    }
};
*/

class Solution {
public:
    Node* copyRandomList(Node* head) {
        if (head == NULL)
            return NULL;
        Node* curr = head;

        //1st step to make copy of all the nodes
        while(curr)
        {
            Node* copy = new Node(curr->val);
            copy->next = curr->next;
            curr->next = copy;
            curr = copy->next;
        }

        //2nd step to add random pointer 

        curr = head ;
        while(curr){
            if (curr->random)
                curr->next->random = curr->random->next; 

            if(curr && curr->next)
                curr = curr->next->next;
        }

        //3rd step seperate the list 

        curr = head;
        Node* head2 = head->next;
        while(curr)
        {
            Node* copy = curr->next;
            curr->next = copy->next;
            curr = copy->next;
            if(curr && copy->next)
                copy->next = curr->next;
        }

        return head2;
    }
};
