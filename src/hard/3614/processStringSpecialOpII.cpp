#include <string>
using namespace std;

class Solution {
public:
  char processStr(const string &s, long k) {
    if (k < 0)
      return '.';

    long n = 0;
    for (const char &c : s)
      switch (c) {
      case '#':
        n *= 2;
        break;
      case '*':
        if (n)
          n--;
        break;
      case '%':
        break;
      default:
        n++;
      }

    if (k >= n)
      return '.';

    for (int i = s.length() - 1; i >= 0; i--) {
      char c = s[i];
      switch (c) {
      case '*':
        n++;
        break;
      case '#':
        n /= 2;
        if (k >= n)
          k -= n;
        break;
      case '%':
        k = n - 1 - k;
        break;
      default:
        if (--n == k)
          return c;
      }
    }

    return '.';
  }
};
