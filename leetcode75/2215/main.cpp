#include <unordered_set>
#include <vector>
using namespace std;

class Solution {
  public:
    vector<vector<int>> findDifference(vector<int> &nums1, vector<int> &nums2) {
        // create a set of nums
        unordered_set<int> s1(nums1.begin(), nums1.end());

        // create a set of nums
        unordered_set<int> s2(nums2.begin(), nums2.end());

        vector<vector<int>> ans;

        unordered_set<int> a1;
        unordered_set<int> a2;

        for (int i = 0; i < nums1.size(); i++) {
            if (s2.find(nums1[i]) == s2.end()) {
                a1.insert(nums1[i]);
            }
        }

        for (int i = 0; i < nums2.size(); i++) {
            if (s1.find(nums2[i]) == s1.end()) {
                a2.insert(nums2[i]);
            }
        }

        ans.push_back(vector<int>(a1.begin(), a1.end()));
        ans.push_back(vector<int>(a2.begin(), a2.end()));

        return ans;
    }
};
