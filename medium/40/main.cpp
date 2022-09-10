#include <bits/stdc++.h>
using namespace std;

class Solution {
public:
  int target;
  vector<vector<int>> result;
  void backtracking(vector<int> &candidates, vector<int> &current, int index,
                    int current_sum) {
    if (index == candidates.size()) {
      // we should check the current_sum
      if (current_sum == target) {
        // copy
        this->result.push_back(current);
      }
      return;
    }
    if (current_sum == target) {
      this->result.push_back(current);
      return;
    }
    if (current_sum > target) {
      return;
    }
    int pre = 0;
    for (int i = index; i < candidates.size(); i++) {
      if (candidates[i] == pre) {
        continue;
      }
      current.push_back(candidates[i]);
      backtracking(candidates, current, i + 1, current_sum + candidates[i]);
      current.pop_back();
      pre = candidates[i];
    }
  }
  vector<vector<int>> combinationSum2(vector<int> &candidates, int target) {
    this->target = target;
    sort(candidates.begin(), candidates.end());
    auto current = vector<int>();
    backtracking(candidates, current, 0, 0);
    return result;
  }
};
