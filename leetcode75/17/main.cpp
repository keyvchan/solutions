#include <iostream>
#include <unordered_map>
#include <vector>
using namespace std;

class Solution {
  public:
    vector<string> ans;

    unordered_map<char, string> phone = {
        {'2', "abc"}, {'3', "def"},  {'4', "ghi"}, {'5', "jkl"},
        {'6', "mno"}, {'7', "pqrs"}, {'8', "tuv"}, {'9', "wxyz"}};

    void dfs(string &digits, int index, string &path) {
        if (index == digits.size()) {
            ans.push_back(path);
            return;
        }

        // iter through the next digit
        auto digit = digits[index];
        for (char letter : phone[digit]) {
            path.push_back(letter);
            dfs(digits, index + 1, path);
            path.pop_back();
        }
    };

    vector<string> letterCombinations(string digits) {
        if (digits.empty()) {
            return {};
        }
        string path;
        dfs(digits, 0, path);
        return this->ans;
    }
};
int main() {
    Solution s;
    auto res = s.letterCombinations("23");
    for (auto r : res) {
        cout << r << endl;
    }

    return 0;
}
