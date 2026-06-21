#include <algorithm>
#include <vector>
using namespace std;

class Solution {
public:
  int maxIceCream(vector<int> &costs, int coins) {
    sort(costs.begin(), costs.end());
    int maxIceCreamPurchase = 0;

    for (int &cost : costs) {
      if (cost <= coins) {
        coins -= cost;
        maxIceCreamPurchase++;
      }
    }

    return maxIceCreamPurchase;
  }
};
