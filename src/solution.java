import java.util.*;
import java.util.stream.Collectors;

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

    //Q54:螺旋矩阵
    public List<Integer> spiralOrder(int[][] matrix) {
        List<Integer> list = new ArrayList<>();
        if(matrix==null||matrix.length==0) return list;
        int len_y = matrix.length;
        int len_x = matrix[0].length;
        int top = 0, bottom = len_y-1, left = 0, right = len_x-1;
        while(top<=bottom&&left<=right){
            if(top==bottom){
                for(int i=left;i<=right;i++)
                    list.add(matrix[top][i]);
            }else if(left==right){
                for(int i=top;i<=bottom;i++)
                    list.add(matrix[i][left]);
            }else{
                for(int i=left;i<=right;i++)
                    list.add(matrix[top][i]);
                for(int i=top+1;i<=bottom;i++)
                    list.add(matrix[i][right]);
                for(int i=right-1;i>=left;i--)
                    list.add(matrix[bottom][i]);
                for(int i=bottom-1;i>top;i--)
                    list.add(matrix[i][left]);
            }
            top++;bottom--;left++;right--;
        }
        return list;
    }

    //Q89:格雷编码
    public List<Integer> grayCode(int n) {
        List<Integer> result = new ArrayList<>();
        if(n==0){
            result.add(0);
            return result;
        }
        for(int i=0; i<1<<n; i++){
            result.add(i^i>>1);
        }
        return result;
    }

    //多个成对筷子，取出其中单独一只
    public int kuaizi(int[] n){
        int tmp = 0;
        for(int i=0;i<n.length;i++){
            tmp ^= n[i];
        }
        return tmp;
    }

    //Q970：强整数
    public List<Integer> powerfulIntegers(int x, int y, int bound) {
        Set<Integer> s = new HashSet<>();
        for(int i=0;i<bound;i++){
            int tmp = (int) Math.pow(x, i);
            if(tmp>bound)break;
            for(int j=0;j<=bound;j++){
                int tmp1 = tmp + (int) Math.pow(y,j);
                if(tmp1>bound)break;
                else s.add(tmp1);
            }
        }
        List<Integer> ret = new ArrayList<>(s);
        return ret;
    }

    //Q15:三数之和
    public List<List<Integer>> threeSum(int[] nums) {
        List<List<Integer>> ret = new ArrayList<>();
        Arrays.sort(nums);
        for(int i=0;i<nums.length;i++){
            if(i!=0&&nums[i]==nums[i-1])continue;
            if(nums[i]>0)break;
            int l = i + 1;
            int r = nums.length - 1;
            while(l<r){
                if(nums[r]<0)break;
                int s = nums[i] + nums[l] + nums[r];
                if(s==0){
                    List<Integer> tmp = new ArrayList<>();
                    tmp.add(nums[i]);
                    tmp.add(nums[l]);
                    tmp.add(nums[r]);
                    ret.add(tmp);
                    int numL = nums[l];
                    while(numL == nums[l]){
                        l++;
                        if(l>=r)break;
                    }
                    int numR = nums[r];
                    while(numR == nums[r]){
                        r--;
                        if(l>=r)break;
                    }
                }else if(s>0)
                    r--;
                else
                    l++;
            }
        }
        return ret;
    }

    //Q55:跳跃游戏
    public boolean canJump(int[] nums) {
        int n = 1;
        for(int i=nums.length-2;i>=0;i--){
            if(nums[i]>=n) n=1;
            else n++;
            if(i==0&&n!=1) return false;
        }
        return true;
    }
}
