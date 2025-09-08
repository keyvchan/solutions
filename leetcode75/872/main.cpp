#include <iostream>
#include <vector>

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
    void dfs(TreeNode *root, vector<int> &leaves) {

        // check if root is leaf
        if (root == nullptr) {
            return;
        }

        cout << "visit: " << root->val << endl;
        // if left and right both true, add it to leaf
        if (root->left == nullptr && root->right == nullptr) {
            cout << "append: " << root->val << endl;
            leaves.push_back(root->val);
            return;
        }
        dfs(root->left, leaves);
        dfs(root->right, leaves);

        return;
    }

    bool leafSimilar(TreeNode *root1, TreeNode *root2) {
        vector<int> leaves1;
        vector<int> leaves2;

        dfs(root1, leaves1);
        dfs(root2, leaves2);

        for (int i = 0; i < leaves1.size(); i++) {
            cout << leaves1[i] << " ";
        }
        cout << endl;
        for (int i = 0; i < leaves2.size(); i++) {
            cout << leaves2[i] << " ";
        }
        cout << endl;

        return std::equal(leaves1.begin(), leaves1.end(), leaves2.begin(),
                          leaves2.end());
    }
};

int main() {
    Solution s;

    TreeNode *root1 = new TreeNode(3);
    root1->left = new TreeNode(5);
    root1->right = new TreeNode(1);

    TreeNode *root2 = new TreeNode(3);
    root2->left = new TreeNode(1);
    root2->right = new TreeNode(5);

    bool res2 = s.leafSimilar(root1, root2);

    cout << res2 << endl;

    TreeNode *root3 = new TreeNode(3);
    root3->left = new TreeNode(5);
    root3->right = new TreeNode(1);
    root3->left->left = new TreeNode(6);
    root3->left->right = new TreeNode(2);
    root3->right->left = new TreeNode(9);
    root3->right->right = new TreeNode(8);
    root3->left->right->left = new TreeNode(7);
    root3->left->right->right = new TreeNode(4);

    TreeNode *root4 = new TreeNode(3);
    root4->left = new TreeNode(5);
    root4->right = new TreeNode(1);
    root4->left->left = new TreeNode(6);
    root4->left->right = new TreeNode(2);
    root4->right->left = new TreeNode(8);
    root4->right->right = new TreeNode(9);
    root4->left->right->left = new TreeNode(4);
    root4->left->right->right = new TreeNode(7);

    bool res3 = s.leafSimilar(root3, root4);

    cout << res3 << endl;

    return 0;
}
