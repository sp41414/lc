#include <vector>
using namespace std;

class Solution {
public:
  int largestAltitude(vector<int> &gain) {
    int currentSum = 0;
    int highest = 0;

    for (int &v : gain) {
      currentSum += v;
      highest = max(highest, currentSum);
    }

    return highest;
  }
};
