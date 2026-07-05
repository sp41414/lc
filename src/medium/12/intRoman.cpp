#include <string>
#include <vector>
using namespace std;

class Solution
{
private:
    inline static const vector<int> VALUES = {
        1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1};
    inline static const vector<string> SYMBOLS = {
        "M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"};

public:
    string intToRoman(int num)
    {
        string result;
        for (size_t i = 0; i < VALUES.size(); ++i)
        {
            while (num >= VALUES[i])
            {
                num -= VALUES[i];
                result.append(SYMBOLS[i]);
            }
        }
        return result;
    }
};
