#include <algorithm>
#include <unordered_set>
#include <vector>
using std::vector;

class Solution {
public:
  vector<vector<int>> findDifference(vector<int> &nums1, vector<int> &nums2) {
    // create two sets
    auto set1 = std::unordered_set<int>(nums1.begin(), nums1.end());
    auto set2 = std::unordered_set<int>(nums2.begin(), nums2.end());

    // create two vectors as results
    vector<int> res1, res2;

    for (auto &num : set1) {
      if (set2.find(num) == set2.end()) {
        res1.push_back(num);
      }
    }

    for (auto &num : set2) {
      if (set1.find(num) == set1.end()) {
        res2.push_back(num);
      }
    }

    return {res1, res2};
  }
};
