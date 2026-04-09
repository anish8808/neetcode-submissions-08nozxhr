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
    int maxHeight(TreeNode* root)
    {
        if(root==NULL)
            return 0;

        int left = maxHeight(root->left);
        int right = maxHeight(root->right);

        return 1 + max(left , right);
    }
    bool isBalanced(TreeNode* root) {
        if (root ==NULL)
            return  true;
       bool left =  isBalanced(root->left);
       bool right = isBalanced(root->right);
        int leftH = maxHeight(root->left);
        int rightH = maxHeight(root->right);

        if (abs(leftH - rightH)>1)
            return false;

        return left && right ;
        
    }
};
