#include <cassert>
#include <iostream>
#include <string>
#include <vector>
using namespace std;

class Solution {
  public:
    bool isSubsequence(string s, string t) {
        int i = 0;
        int j = 0;

        while (i < s.size()) {
            // take a char of s, find corrsponding char in t

            auto pos = t.find_first_of(s[i], j);
            if (pos == string::npos) {
                return false;
            }
            j = pos;
            j++;

            //
            i++;
        }

        return true;
    }
};

int main() {
    Solution s;
    // assert(s.isSubsequence("abc", "ahbgdc"));
    // assert(!s.isSubsequence("axc", "ahbgdc"));
    assert(!s.isSubsequence("aaaaaa", "bbaaaa"));
}
