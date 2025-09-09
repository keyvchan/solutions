#include <iostream>
#include <queue>
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
    vector<int> rightSideView(TreeNode *root) {
        if (root == nullptr) {
            return {};
        }
        vector<int> result;
        // we should have a queue, to store a level of nodes
        deque<TreeNode *> level;
        deque<TreeNode *> nextLevel;

        level.push_back(root);

        while (!level.empty()) {
            auto rightest = level.back();
            result.push_back(rightest->val);

            // put nodes in nextLevel
            while (!level.empty()) {
                // pop_back, store the first nodes
                auto front = level.front();
                if (front->left != nullptr) {
                    nextLevel.push_back(front->left);
                }
                if (front->right != nullptr) {
                    nextLevel.push_back(front->right);
                }
                level.pop_front();
            }

            // emplace level with nextLevel
            level = nextLevel;
            nextLevel.clear();
        }

        return result;
    }
};

int main() {
    TreeNode *root = new TreeNode(1);
    root->left = new TreeNode(2);
    root->right = new TreeNode(3);
    root->left->right = new TreeNode(5);
    Solution s;
    auto result = s.rightSideView(root);
    for (int i = 0; i < result.size(); i++) {
        cout << result[i] << " ";
    }
    cout << endl;
    return 0;
}
