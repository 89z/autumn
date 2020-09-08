#include <functional>
#include <numeric>
int s = std::accumulate(x.begin(), x.end(), 0, std::plus<int>());
