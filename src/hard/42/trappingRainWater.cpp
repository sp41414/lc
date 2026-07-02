#include <vector>
using namespace std;

class Solution
{
public:
    int trap(vector<int> &height)
    {
        int n = height.size();
        if (n <= 2)
            return 0;

        int max_height = 0, max_height_idx = 0;
        for (int i = 0; i < n; i++)
        {
            if (height[i] > max_height)
            {
                max_height_idx = i;
                max_height = height[i];
            }
        }

        int total = 0;

        int curr_max_height = 0;
        for (int i = 0; i < max_height_idx; i++)
        {
            if (height[i] > curr_max_height)
            {
                curr_max_height = height[i];
                continue;
            }
            else
            {
                total += curr_max_height - height[i];
            }
        }

        curr_max_height = 0;
        for (int i = n - 1; i > max_height_idx; i--)
        {
            if (height[i] > curr_max_height)
            {
                curr_max_height = height[i];
                continue;
            }
            else
            {
                total += curr_max_height - height[i];
            }
        }

        return total;
    }
};
