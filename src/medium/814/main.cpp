#include <bits/stdc++.h>

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
  bool preorder(TreeNode *root) {
    // postorder
    if (root == nullptr) {
      return false;
    }

    auto left_result = preorder(root->left);
    auto right_result = preorder(root->right);
    // prune the false one
    if (left_result == false) {
      root->left = nullptr;
    }
    if (right_result == false) {
      root->right = nullptr;
    }

    if (root->left == nullptr && root->right == nullptr) {
      return root->val == 1;
    } else {
      return true;
    }
  }
  TreeNode *pruneTree(TreeNode *root) {
    // post order
    if (preorder(root) == false) {
      return nullptr;
    }
    return root;
  }
};
