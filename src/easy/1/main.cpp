#include <bits/stdc++.h>

using namespace std;

class Solution {
public:
  vector<int> twoSum(vector<int> &nums, int target) {

    // create a hash map
    auto nums_map = unordered_map<int, int>();

    for (auto i = 0; i < nums.size(); i++) {
      auto finded_index = nums_map.find(target - nums[i]);

      if (finded_index != nums_map.end()) {
        return {finded_index->second, i};
      } else {
        nums_map[nums[i]] = i;
      }
    }
    return {};
  }
};
