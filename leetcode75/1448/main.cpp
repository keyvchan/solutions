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
    int count = 1;
    void dfs(TreeNode *root, int max) {
        if (root == nullptr) {
            return;
        }

        // if current node is greater than the max, increase count
        if (root->val >= max) {
            count++;
            max = root->val;
        }

        dfs(root->left, max);
        dfs(root->right, max);
    }

    int goodNodes(TreeNode *root) {
        // matain a max value of a path
        dfs(root->left, root->val);
        dfs(root->right, root->val);

        return count;
    }
};

int main() {
    TreeNode *root = new TreeNode(3);
    root->left = new TreeNode(1);
    root->right = new TreeNode(4);
    root->left->left = new TreeNode(3);
    root->left->right = new TreeNode(1);
    root->right->right = new TreeNode(5);

    TreeNode *root2 = new TreeNode(3);
    root2->left = new TreeNode(3);
    root2->left->left = new TreeNode(4);
    root2->left->right = new TreeNode(2);

    Solution s;
    cout << s.goodNodes(root) << endl;

    Solution s2;
    cout << s2.goodNodes(root2) << endl;

    return 0;
}
