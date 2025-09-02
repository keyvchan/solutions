#include <iostream>
#include <string>
#include <unordered_set>
using namespace std;

class Solution {

  public:
    unordered_map<char, bool> vowels{
        {'a', true}, {'e', true}, {'i', true}, {'o', true}, {'u', true},
    };

    int maxVowels(string s, int k) {
        // create a sliding window

        int left = 0;
        int right = k - 1;

        // res for current window
        int currentRest = 0;

        // res for max
        int maxRes = 0;

        for (int i = 0; i < k; i++) {
            if (vowels[s[i]]) {
                currentRest++;
            }
        }
        maxRes = currentRest;

        while (right + 1 < s.size()) {
            if (vowels[s[left]]) {
                currentRest--;
            }

            if (vowels[s[right + 1]]) {
                currentRest++;
            }

            maxRes = max(maxRes, currentRest);
            left++;
            right++;
        }

        return maxRes;
    }
};

int main() {
    Solution s;
    cout << s.maxVowels("abciiidef", 3) << endl;
    return 0;
}
