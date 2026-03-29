/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode() : val(0), left(nullptr), right(nullptr) {}
 *     TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
 *     TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
 * };
 */

class Solution {
public:
    bool isSameTree(TreeNode* p , TreeNode* q)
    {
        if(p==NULL && q==NULL)
            return true;
        if(p!=NULL && q== NULL)
            return false;
        if(p==NULL && q!=NULL)
            return false;

        bool left = true;
        bool right = true;

        if(p->val == q->val)
        {
           left = left &&  isSameTree(p->left , q->left);
           right = right &&  isSameTree(p->right , q->right);
        }else {
            return false;
        }

        return left && right ;

    }
    bool isSubtree(TreeNode* root, TreeNode* subRoot) {
        
        bool ans = false ;
         if(root==NULL && subRoot==NULL)
            return true;
        if(root!=NULL && subRoot== NULL)
            return false;
        if(root==NULL && subRoot!=NULL)
            return false;
        
        if(root->val == subRoot->val)
        {
            ans = ans || isSameTree(root , subRoot);
        }

        bool left = isSubtree(root->left , subRoot);
        bool right = isSubtree(root->right , subRoot);

        return ans || left || right;
    }
};
