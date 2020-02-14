//
// Created by polar on 2020/2/14.
//
#include <vector>
#include <iostream>

using namespace std;

class Solution {
public:
    vector<vector<int>> threeSum(vector<int> &nums) {
        int numOfN = nums.size(), sum;
        if (numOfN <= 2)
            return {};
        vector<vector<int>> ans;
        sort(nums.begin(), nums.end()); // if you'd not change the origin vector, copy it to a new one before sort
        for (int i = 0; i <= numOfN - 3; i++)
        {
            //jump the same first value to avoid same result since it's sorted
            if (i >= 1 && nums[i] == nums[i - 1])
                continue;

            for (int l = i + 1, r = numOfN - 1; l < r;)
            {
                sum = nums[i] + nums[l] + nums[r];

                //since it is sorted
                if (sum < 0) {l++; continue;}
                if (sum > 0) {r--; continue;}
                ans.push_back({nums[i], nums[l], nums[r]});

                //l and r go to next pos with different value
                l++;
                while (l < r && nums[l] == nums[l - 1])
                    l++;
                r--;
                while (l < r && nums[r] == nums[r + 1])
                    r--;
            }

        }
        return ans;
    }
};

int main()
{
    vector<int> nums = {-1, 0, 1, 2, -1, -4};
    Solution s;
    auto ret = s.threeSum(nums);
    for(auto nlist : ret){
        for(auto n : nlist){
            cout << n << " ";
        }
        cout << endl;
    }
    return 0;
}