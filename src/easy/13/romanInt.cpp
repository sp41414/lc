#include <string>
#include <map>
using namespace std;

class Solution
{
private:
    inline static const map<char, int> ROMAN_NUMERALS = {
        {'I', 1},
        {'V', 5},
        {'X', 10},
        {'L', 50},
        {'C', 100},
        {'D', 500},
        {'M', 1000},
    };

public:
    int romanToInt(string s)
    {
        int total = 0;
        int n = s.size();
        for (int i = 0; i < n - 1; i++)
        {
            if (ROMAN_NUMERALS.at(s[i]) < ROMAN_NUMERALS.at(s[i + 1]))
            {
                total -= ROMAN_NUMERALS.at(s[i]);
            }
            else
            {
                total += ROMAN_NUMERALS.at(s[i]);
            }
        }
        return total + ROMAN_NUMERALS.at(s[n - 1]);
    }
};
