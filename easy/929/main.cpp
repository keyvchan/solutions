#include <string>
#include <string_view>
#include <unordered_set>
#include <vector>
using std::string;
using std::string_view;
using std::unordered_set;
using std::vector;

class Solution {
public:
  int numUniqueEmails(vector<string> &emails) {
    // split the email into local and domain
    //
    auto unique_emails = unordered_set<string>();

    auto count = 0;
    for (auto &email : emails) {
      // split the email into local and domain

      auto email_view = string_view(email);

      auto at_pos = email_view.find('@');

      auto local = email_view.substr(0, at_pos);
      auto domain = email_view.substr(at_pos + 1);

      // remove the dots
      auto new_local = string();
      new_local.reserve(local.size() + 1 + domain.size());
      for (auto c : local) {
        if (c == '.') {
          continue;
        }
        if (c == '+') {
          break;
        }
        new_local += c;
      }

      auto new_email = new_local + "@" + domain.data();

      // check if the email is unique
      if (unique_emails.find(new_email) != unique_emails.end()) {
        continue;
      }

      count++;
      unique_emails.insert(new_email);
    }

    return count;
  }
};
