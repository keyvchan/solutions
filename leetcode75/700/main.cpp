
#include <iostream>
using namespace std;

struct TreeNode {
    int val;
    TreeNode *left;
    TreeNode *right;
    TreeNode() : val(0), left(nullptr), right(nullptr) {}
    TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
    TreeNode(int x, TreeNode *left, TreeNode *right)
        : val(x), left(left), right(right) {}
};

class Solution {
  public:
    TreeNode *matched = nullptr;
    void dfs(TreeNode *root, int val) {
        if (root == nullptr) {
            return;
        }

        // check root
        if (root->val == val) {
            // equal
            matched = root;
            return;
        }

        // check val to determe search left or right
        if (root->val < val) {
            // root.val < val, should search right
            dfs(root->right, val);
        } else {
            // search left
            dfs(root->left, val);
        }
        return;
    }
    TreeNode *searchBST(TreeNode *root, int val) {
        dfs(root, val);
        return matched;
    }
};

int main() {
    TreeNode *root = new TreeNode(4);
    root->left = new TreeNode(2);
    root->right = new TreeNode(7);
    root->left->left = new TreeNode(1);
    root->left->right = new TreeNode(3);
    root->right->left = new TreeNode(6);
    root->right->right = new TreeNode(9);

    Solution s;
    TreeNode *result = s.searchBST(root, 2);

    cout << result->val << endl;

    return 0;
}
