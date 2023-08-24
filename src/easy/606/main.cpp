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
  void preorder(TreeNode *root, string &result) {
    if (root == nullptr) {
      return;
    }

    // concat
    result += to_string(root->val);

    if (root->left == nullptr && root->right != nullptr) {
      result += "()";
    }
    if (!(root->left == nullptr)) {
      result += "(";
    }
    preorder(root->left, result);
    if (!(root->left == nullptr)) {
      result += ")";
    }
    if (!(root->right == nullptr)) {
      result += "(";
    }
    preorder(root->right, result);
    if (!(root->right == nullptr)) {
      result += ")";
    }
  }

  string tree2str(TreeNode *root) {
    // preorder
    string result = "";
    preorder(root, result);

    return result;
  }
};
