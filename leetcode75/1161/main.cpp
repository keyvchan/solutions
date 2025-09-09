#include <deque>
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
    int maxLevelSum(TreeNode *root) {
        if (root == nullptr) {
            return 1;
        }

        int curLevelCount = 0;
        int maxVal = root->val;
        int maxLevel = 1;
        deque<TreeNode *> currentLevel;
        deque<TreeNode *> nextLevel;

        currentLevel.push_back(root);

        while (!currentLevel.empty()) {
            // take node from front anc cal the sum
            int sum = 0;
            while (!currentLevel.empty()) {
                TreeNode *front = currentLevel.front();
                sum += front->val;
                if (front->left != nullptr) {
                    nextLevel.push_back(front->left);
                }
                if (front->right != nullptr) {
                    nextLevel.push_back(front->right);
                }
                currentLevel.pop_front();
            }
            curLevelCount++;

            if (sum > maxVal) {
                maxVal = sum;
                maxLevel = curLevelCount;
            }
            currentLevel = nextLevel;
            nextLevel.clear();
        }
        return maxLevel;
    }
};
