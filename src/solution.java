import java.util.HashMap;
import java.util.Map;

public class solution {

    //Q1:两数之和
    //暴力法
    public int[] twoSum(int[] nums, int target) {
        for (int i = 0; i < nums.length; i++) {
            for (int j = i + 1; j < nums.length; j++) {
                if (nums[j] == target - nums[i]) {
                    return new int[] { i, j };
                }
            }
        }
        throw new IllegalArgumentException("No two sum solution");
    }
    //哈希索引
    public int[] twoSum2(int[] nums, int target) {
        Map<Integer, Integer> map = new HashMap<>();
        for (int i = 0; i < nums.length; i++) {
            int complement = target - nums[i];
            if (map.containsKey(complement)) {
                return new int[] { map.get(complement), i };
            }
            map.put(nums[i], i);
        }
        throw new IllegalArgumentException("No two sum solution");
    }

    //Q2:盛最多水的容器
    public int maxArea(int[] height) {
        int maxarea = 0, l = 0, r = height.length-1;
        while (l<r){
            maxarea = Math.max(maxarea, (r-l)*(Math.min(height[l], height[r])));
            if(height[l]< height[r])
                l++;
            else
                r--;
        }
        return maxarea;
    }
}
