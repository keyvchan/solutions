#include <vector>
using std::vector;

class Solution {
public:
  int largestAltitude(vector<int> &gain) {
    auto altitude = 0;
    auto max_altitude = 0;
    for (auto &g : gain) {
      altitude += g;
      if (altitude > max_altitude) {
        max_altitude = altitude;
      }
    }
    return max_altitude;
  }
};
