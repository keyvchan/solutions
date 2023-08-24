#include <bits/stdc++.h>
using namespace std;

class Solution {
public:
  vector<int> successfulPairs(vector<int> &spells, vector<int> &potions,
                              long long success) {
    vector<int> ans;

    // sort the potions
    sort(potions.begin(), potions.end());

    // brute force
    for (int i = 0; i < spells.size(); i++) {
      int counts = 0;
      // calculate the success / spells[i]
      long long successDiv = success / spells[i];
      long long reminder = success % spells[i];

      vector<int>::iterator bound;

      if (reminder != 0) {
        // we can safely find the upper bound
        bound = upper_bound(potions.begin(), potions.end(), successDiv);
      } else {
        // we need to find the lower bound
        bound = lower_bound(potions.begin(), potions.end(), successDiv);
      }
      counts = distance(bound, potions.end());

      ans.push_back(counts);
    }

    return ans;
  }
};
int main() {
  Solution s;
  vector<int> spells = {3, 1, 2};
  vector<int> potions = {8, 5, 8};
  long long success = 10;
  vector<int> ans = s.successfulPairs(spells, potions, success);
  for (auto i : ans) {
    cout << i << " ";
  }
  return 0;
}
