#include <bits/stdc++.h>
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
  void inorder(TreeNode *root, vector<int> &result) {
    if (root == nullptr) {
      return;
    }
    inorder(root->left, result);
    result.push_back(root->val);
    inorder(root->right, result);
  }
  vector<int> inorderTraversal(TreeNode *root) {
    auto result = vector<int>();
    inorder(root, result);
    return result;
  }
};