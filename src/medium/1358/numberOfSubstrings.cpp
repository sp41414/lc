#include <string>
#include <algorithm>
using namespace std;

/*
    it basically tracks the last seen index for each a, b, and c
    and looks backwards by finding the min of the last seen indices
*/
class Solution
{
public:
    int numberOfSubstrings(string &s)
    {
        int last_a = -1, last_b = -1, last_c = -1;
        int result = 0;

        for (int i = 0; i < s.size(); i++)
        {
            switch (s[i])
            {
            case 'a':
                last_a = i;
                break;
            case 'b':
                last_b = i;
                break;
            case 'c':
                last_c = i;
                break;
            }

            if (last_a != -1 && last_b != -1 && last_c != -1)
            {
                result += min({last_a, last_b, last_c}) + 1;
            }
        }

        return result;
    }
};
