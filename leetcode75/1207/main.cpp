#include <unordered_map>
#include <unordered_set>
#include <vector>
using namespace std;

class Solution {
  public:
    bool uniqueOccurrences(vector<int> &arr) {
        unordered_map<int, int> m;
        for (auto i : arr) {
            m[i]++;
        }

        unordered_set<int> occ;
        for (auto it : m) {
            if (occ.find(it.second) != occ.end()) {
                return false;
            }
            occ.insert(it.second);
        }
        return true;
    }
};
