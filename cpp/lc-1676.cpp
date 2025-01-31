#include <iostream>
#include <map>

using namespace std;

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
    TreeNode* lowestCommonAncestor(TreeNode* root, vector<TreeNode*> &nodes) {
        map<int, bool> dict;   
        for (const auto &node : nodes) {
            cout << node->val << endl;
            dict[node->val] = true;
        }
        TreeNode* targetNode = NULL;
        search(root, dict, targetNode);
        return targetNode;
    }
    
    int search(TreeNode* node, map<int, bool> &nodeDict, TreeNode* &targetNode) {
        if (node == NULL) {
            return 0;
        }
        int res = 0;
        if (nodeDict.find(node->val) != nodeDict.end()) {
            ++res;
        }
        res += search(node->left, nodeDict, targetNode) + search(node->right, nodeDict, targetNode);
        if (res == nodeDict.size() && targetNode == NULL) {
            targetNode = node;
        }
        return res;
    }
};
