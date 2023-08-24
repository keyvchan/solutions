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
  int total;
  void dfs(TreeNode *root, unordered_map<int, int> &visited) {
    if (visited.find(root->val) != visited.end()) {
      // we find a key
      visited[root->val]--;
      if (visited[root->val] == 0) {
        visited.erase(root->val);
      }
    } else {
      visited[root->val]++;
    }

    if (root->left != nullptr) {
      dfs(root->left, visited);
      if (visited.find(root->left->val) == visited.end()) {
        visited[root->left->val]++;
      } else {
        visited[root->left->val]--;
      }
      if (visited[root->left->val] == 0) {
        visited.erase(root->left->val);
      }
    }
    if (root->right != nullptr) {
      dfs(root->right, visited);
      if (visited.find(root->right->val) == visited.end()) {
        visited[root->right->val]++;
      } else {
        visited[root->right->val]--;
      }
      if (visited[root->right->val] == 0) {
        visited.erase(root->right->val);
      }
    }
    if (root->left == nullptr && root->right == nullptr) {
      // check visited
      if (visited.size() <= 1) {
        total++;
      }
      return;
    }
  }
  int pseudoPalindromicPaths(TreeNode *root) {
    // dfs
    auto visited = unordered_map<int, int>();
    dfs(root, visited);
    return total;
  }
};
