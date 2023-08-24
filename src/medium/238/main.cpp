#include <bits/stdc++.h>
using namespace std;

class Solution {
public:
  vector<int> productExceptSelf(vector<int> &nums) {
    // prefix product and suffix product
    auto prefix = vector<int>(nums.size(), 0);
    for (int i = 0; i < nums.size(); i++) {
      if (i == 0) {
        prefix[i] = nums[i];
        continue;
      }
      prefix[i] = nums[i] * prefix[i - 1];
    }
    auto suffix = vector<int>(nums.size(), 0);
    for (int i = nums.size() - 1; i >= 0; i--) {
      if (i == nums.size() - 1) {
        suffix[i] = nums[i];
        continue;
      }
      suffix[i] = nums[i] * suffix[i + 1];
    }

    auto result = vector<int>(nums.size(), 0);
    for (int i = 0; i < nums.size(); i++) {
      if (i == 0) {
        result[i] = suffix[i + 1];
        continue;
      }
      if (i == nums.size() - 1) {
        result[i] = prefix[i - 1];
        continue;
      }
      result[i] = prefix[i - 1] * suffix[i + 1];
    }
    return result;
  }
};
