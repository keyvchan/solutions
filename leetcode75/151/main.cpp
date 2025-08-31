#include <cassert>
#include <iostream>
#include <string>
#include <vector>

using namespace std;

class Solution {
  public:
    vector<string_view> split(string_view s) {
        vector<string_view> result;
        string_view::size_type prev_pos = 0, pos = 0;
        for (; pos < s.size();) {
            if (s[pos] != ' ') {
                // start counting
                prev_pos = pos;
                pos = s.find_first_of(' ', pos);
                if (pos == string_view::npos) {
                    pos = s.size();
                }
                result.push_back(s.substr(prev_pos, pos - prev_pos));
            } else {
                pos++;
            }
        }

        return result;
    }
    string reverseWords(string s) {
        auto words = split(s);

        string result;
        for (int index = words.size() - 1; index >= 0; index--) {
            result += words[index];
            if (index != 0) {
                result += ' ';
            }
        }
        return result;
    }
};

int main() {
    Solution s;
    assert(s.reverseWords("the sky is blue") == "blue is sky the");
    assert(s.reverseWords("  hello world!  ") == "world! hello");
    assert(s.reverseWords("a good   example") == "example good a");
    return 0;
}
