#include <bits/stdc++.h>
using namespace std;

class Solution {
public:
  vector<int> topKFrequent(vector<int> &nums, int k) {
    // use a hash map to count frequency
    auto int_map = unordered_map<int, int>();

    for (auto &n : nums) {
      int_map[n]++;
    }
    // put value into vector
    auto pairs = vector<pair<int, int>>(int_map.begin(), int_map.end());
    sort(pairs.begin(), pairs.end(),
         [](auto &x, auto &y) { return x.second < y.second; });

    auto result = vector<int>();
    for (int i = 0; i < k; i++) {
      result.push_back(pairs[pairs.size() - k - 1].first);
    }
    return result;
  }
};
