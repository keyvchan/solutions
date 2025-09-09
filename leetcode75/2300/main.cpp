#include <iostream>
#include <vector>
using namespace std;

class Solution {
  public:
    vector<int> successfulPairs(vector<int> &spells, vector<int> &potions,
                                long long success) {
        // sort potions
        sort(potions.begin(), potions.end());

        vector<int> result;
        result.reserve(spells.size());

        for (int i = 0; i < spells.size(); i++) {

            long long left = 0;
            long long right = potions.size() - 1;

            long long middle = (left + right) / 2 - 1;

            while (left <= right) {

                middle = left + (right - left) / 2;

                // check middle is what we looking for
                long long product = static_cast<long long>(potions[middle]) *
                                    static_cast<long long>(spells[i]);
                if (product >= success) {
                    right = middle - 1;
                } else {
                    left = middle + 1;
                }
            }

            result.push_back(
                distance(potions.begin() + right + 1, potions.end()));
        }
        return result;
    }
};

int main() {
    vector<int> spells = {3, 1, 2};
    vector<int> potions = {8, 5, 8};
    long long success = 16;
    Solution s;
    auto result = s.successfulPairs(spells, potions, success);

    for (int i : result) {
        cout << i << " ";
    }
    cout << endl;

    vector<int> spells2 = {5, 1, 3};
    vector<int> potions2 = {1, 2, 3, 4, 5};
    long long success2 = 7;
    auto result2 = s.successfulPairs(spells2, potions2, success2);

    for (int i : result2) {
        cout << i << " ";
    }
    cout << endl;
    return 0;
}
