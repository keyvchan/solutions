#include <bits/stdc++.h>
using namespace std;

class Solution {
public:
  int buy_limit;
  vector<int> prices;
  unordered_map<string, int> memo;

  int dfs(int day, bool buyed, int sell_counts) {
    if (day == prices.size()) {
      return 0;
    }
    // serilize the state
    auto state =
        to_string(day) + "," + to_string(buyed) + "," + to_string(sell_counts);
    if (this->memo.find(state) != this->memo.end()) {
      return this->memo[state];
    }
    int max_profit = 0;
    if (buyed) {
      // we already buyed, we can't buy more
      // sell
      max_profit =
          max(max_profit, dfs(day + 1, false, sell_counts + 1) + prices[day]);
      // not sell
      max_profit = max(max_profit, dfs(day + 1, buyed, sell_counts));
    } else {
      // check buy limit
      if (sell_counts < buy_limit) {
        // we can buy more, or not buying
        max_profit =
            max(max_profit, dfs(day + 1, true, sell_counts) - prices[day]);
        max_profit = max(max_profit, dfs(day + 1, buyed, sell_counts));
      } else {
        // we can't buy more
      }
    }
    this->memo[state] = max_profit;
    return max_profit;
  }

  int maxProfit(int k, vector<int> &prices) {
    this->buy_limit = k;
    this->prices = prices;
    // for every day, we have two choices
    return dfs(0, false, 0);
  }
};
