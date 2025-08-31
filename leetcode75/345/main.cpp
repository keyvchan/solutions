#include <algorithm>
#include <cassert>
#include <iostream>
#include <string>
#include <unordered_map>
#include <vector>
using namespace std;

class Solution {
  public:
    string reverseVowels(string s) {
        // collect vowels
        vector<char> vowels;

        unordered_map<char, bool> m;
        m['a'] = 1;
        m['e'] = 1;
        m['i'] = 1;
        m['o'] = 1;
        m['u'] = 1;
        m['A'] = 1;
        m['E'] = 1;
        m['I'] = 1;
        m['O'] = 1;
        m['U'] = 1;

        for (auto &c : s) {
            if (m[c]) {
                vowels.push_back(c);
            }
        }

        auto i = vowels.rbegin();
        for (auto &c : s) {
            if (m[c]) {
                c = *i;
                i++;
            }
        }

        return s;
    }
};

int main() {
    assert(Solution().reverseVowels("hello") == "holle");
    assert(Solution().reverseVowels("leetcode") == "leotcede");
    assert(Solution().reverseVowels("aA") == "Aa");
    assert(Solution().reverseVowels("Aa") == "aA");
    assert(Solution().reverseVowels("a") == "a");
    return 0;
}
