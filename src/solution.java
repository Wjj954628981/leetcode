import util.*;

import java.util.*;

public class solution {

    //Q1:两数之和
    //
    //给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
    //你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。
    //
    //示例:
    //给定 nums = [2, 7, 11, 15], target = 9
    //因为 nums[0] + nums[1] = 2 + 7 = 9
    //所以返回 [0, 1]
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

    //Q11:盛最多水的容器
    //给定 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0)。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
    //
    //说明：你不能倾斜容器，且 n 的值至少为 2。
    //
    //
    //
    //图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。
    //
    //
    //
    //示例:
    //
    //输入: [1,8,6,2,5,4,8,3,7]
    //输出: 49
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
    //给定一个包含 m x n 个元素的矩阵（m 行, n 列），请按照顺时针螺旋顺序，返回矩阵中的所有元素。
    //
    //示例 1:
    //输入:
    //[
    // [ 1, 2, 3 ],
    // [ 4, 5, 6 ],
    // [ 7, 8, 9 ]
    //]
    //输出: [1,2,3,6,9,8,7,4,5]
    //示例 2:
    //输入:
    //[
    //  [1, 2, 3, 4],
    //  [5, 6, 7, 8],
    //  [9,10,11,12]
    //]
    //输出: [1,2,3,4,8,12,11,10,9,5,6,7]
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
    //格雷编码是一个二进制数字系统，在该系统中，两个连续的数值仅有一个位数的差异。
    //给定一个代表编码总位数的非负整数 n，打印其格雷编码序列。格雷编码序列必须以 0 开头。
    //
    //示例 1:
    //
    //输入: 2
    //输出: [0,1,3,2]
    //解释:
    //00 - 0
    //01 - 1
    //11 - 3
    //10 - 2
    //
    //对于给定的 n，其格雷编码序列并不唯一。
    //例如，[0,2,3,1] 也是一个有效的格雷编码序列。
    //00 - 0
    //10 - 2
    //11 - 3
    //01 - 1
    //
    //示例 2:
    //输入: 0
    //输出: [0]
    //解释: 我们定义格雷编码序列必须以 0 开头。
    //     给定编码总位数为 n 的格雷编码序列，其长度为 2n。当 n = 0 时，长度为 20 = 1。
    //     因此，当 n = 0 时，其格雷编码序列为 [0]。
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
    //给定两个正整数 x 和 y，如果某一整数等于 x^i + y^j，其中整数 i >= 0 且 j >= 0，那么我们认为该整数是一个强整数。
    //
    //返回值小于或等于 bound 的所有强整数组成的列表。
    //
    //你可以按任何顺序返回答案。在你的回答中，每个值最多出现一次。
    //
    //
    //
    //示例 1：
    //
    //输入：x = 2, y = 3, bound = 10
    //输出：[2,3,4,5,7,9,10]
    //解释：
    //2 = 2^0 + 3^0
    //3 = 2^1 + 3^0
    //4 = 2^0 + 3^1
    //5 = 2^1 + 3^1
    //7 = 2^2 + 3^1
    //9 = 2^3 + 3^0
    //10 = 2^0 + 3^2
    //示例 2：
    //
    //输入：x = 3, y = 5, bound = 15
    //输出：[2,4,6,8,10,14]
    //
    //
    //提示：
    //
    //1 <= x <= 100
    //1 <= y <= 100
    //0 <= bound <= 10^6
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
    //给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？找出所有满足条件且不重复的三元组。
    //
    //注意：答案中不可以包含重复的三元组。
    //
    //例如, 给定数组 nums = [-1, 0, 1, 2, -1, -4]，
    //
    //满足要求的三元组集合为：
    //[
    //  [-1, 0, 1],
    //  [-1, -1, 2]
    //]
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
    //给定一个非负整数数组，你最初位于数组的第一个位置。
    //
    //数组中的每个元素代表你在该位置可以跳跃的最大长度。
    //
    //判断你是否能够到达最后一个位置。
    //
    //示例 1:
    //
    //输入: [2,3,1,1,4]
    //输出: true
    //解释: 从位置 0 到 1 跳 1 步, 然后跳 3 步到达最后一个位置。
    //示例 2:
    //
    //输入: [3,2,1,0,4]
    //输出: false
    //解释: 无论怎样，你总会到达索引为 3 的位置。但该位置的最大跳跃长度是 0 ， 所以你永远不可能到达最后一个位置。
    public boolean canJump(int[] nums) {
        int n = 1;
        for(int i=nums.length-2;i>=0;i--){
            if(nums[i]>=n) n=1;
            else n++;
            if(i==0&&n!=1) return false;
        }
        return true;
    }

    //Q56:合并区间
    public List<Interval> merge(List<Interval> intervals) {
        if(intervals == null || intervals.size() <= 1){
            return intervals;
        }
        Collections.sort(intervals, new IntervalComparator());
        LinkedList<Interval> ret = new LinkedList<>();
        for (Interval tmp:intervals){
            if(ret.isEmpty() || ret.getLast().end<tmp.start){
                ret.add(tmp);
            }else{
                ret.getLast().end = Math.max(ret.getLast().end, tmp.end);
            }
        }
        return ret;
    }
    private class IntervalComparator implements Comparator<Interval> {
        public int compare(Interval I1, Interval I2) {
            return I1.start-I2.start;
        }
    }

    //Q670:最大交换
    //给定一个非负整数，你至多可以交换一次数字中的任意两位。返回你能得到的最大值。
    //
    //示例 1 :
    //
    //输入: 2736
    //输出: 7236
    //解释: 交换数字2和数字7。
    //示例 2 :
    //
    //输入: 9973
    //输出: 9973
    //解释: 不需要交换。
    //注意:
    //
    //给定数字的范围是 [0, 108]
    public int maximumSwap(int num) {
        char[] digits = Integer.toString(num).toCharArray();
        int[] buckets = new int[10];
        for (int i = 0; i < digits.length; i++) {
            buckets[digits[i] - '0'] = i;
        }
        for (int i = 0; i < digits.length; i++) {
            for (int k = 9; k > digits[i] - '0'; k--) {
                if (buckets[k] > i) {
                    char tmp = digits[i];
                    digits[i] = digits[buckets[k]];
                    digits[buckets[k]] = tmp;
                    return Integer.valueOf(new String(digits));
                }
            }
        }
        return num;
    }


    //for:addTwoNumbers
    //对一个链表进行+1操作
    public ListNode addOne(ListNode node){
        int sum = 0, carry = 1;
        ListNode root = node;
        while(true){
            sum = node.val + carry;
            carry = sum >= 10?1:0;
            node.val = sum % 10;
            if(node.next == null)break;
            node = node.next;
        }
        if(carry == 1)node.next = new ListNode(1);
        return root;
    }

    //Q2:两数相加
    //给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。
    //如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
    //您可以假设除了数字 0 之外，这两个数都不会以 0 开头。
    //
    //示例：
    //输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
    //输出：7 -> 0 -> 8
    //原因：342 + 465 = 807
    public ListNode addTwoNumbers(ListNode l1, ListNode l2) {
        int sum, carry = 0;
        ListNode root = l1;
        while(true){
            sum = l1.val + l2.val + carry;
            carry = sum >= 10?1:0;
            l1.val = sum % 10;
            if(l1.next == null && l2.next == null){
                if(carry == 1){
                    ListNode leaf = new ListNode(1);
                    l1.next = leaf;
                }
                break;
            }else if(l1.next == null){
                if(carry == 1)
                    l1.next = addOne(l2.next);
                else
                    l1.next = l2.next;
                break;
            }else if(l2.next == null){
                if(carry == 1)
                    l1.next = addOne(l1.next);
                break;
            }else{
                l1 = l1.next;
                l2 = l2.next;
            }
        }
        return root;
    }

    //Q3:无重复字符的最长子串
    //给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
    //
    //示例 1:
    //
    //输入: "abcabcbb"
    //输出: 3
    //解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
    //示例 2:
    //
    //输入: "bbbbb"
    //输出: 1
    //解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
    //示例 3:
    //
    //输入: "pwwkew"
    //输出: 3
    //解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
    //请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
    public int lengthOfLongestSubstring(String s) {
        if(s.isEmpty())return 0;
        Map<Character, Integer> map = new HashMap<>();
        int left = -1, ret = 0;
        char val;
        for(int i=0;i<s.length();i++){
            val = s.charAt(i);
            if(map.get(val) != null)
                left = Math.max(left, map.get(val));
            ret = Math.max(ret, i-left);
            map.put(val, i);
//            System.out.println(i + "===" + val + "===" + ret + "===" + left + "===" + map.toString());
        }
        return ret;
    }

    public int lengthOfLongestSubstring2(String s) {
        int max = 0;
        int length = 0;
        char val;
        Map<Character, Integer> map = new HashMap<>();
        for (int i = 0; i < s.length(); ++i) {
            val = s.charAt(i);
            Integer index = map.get(val);
            if (index == null) {
                ++length;
            } else if (index < i - length) {
                ++length;
            } else {
                max = Math.max(length, max);
                length = i-index;
            }
            map.put(val, i);
        }
        max = Math.max(length, max);
        return max;
    }

    //Q4:寻找两个有序数组的中位数
    //给定两个大小为 m 和 n 的有序数组 nums1 和 nums2。
    //
    //请你找出这两个有序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。
    //
    //你可以假设 nums1 和 nums2 不会同时为空。
    //
    //示例 1:
    //
    //nums1 = [1, 3]
    //nums2 = [2]
    //
    //则中位数是 2.0
    //示例 2:
    //
    //nums1 = [1, 2]
    //nums2 = [3, 4]
    //
    //则中位数是 (2 + 3)/2 = 2.5
    public double findMedianSortedArrays(int[] nums1, int[] nums2) {
        int length1 = nums1.length, length2 = nums2.length, pos1 = 0, pos2 = 0, k = 0;
        int[] ret = new int[length1 + length2];
        while(pos1 < length1 && pos2 < length2){
            if(nums1[pos1] <= nums2[pos2]){
                ret[k++] = nums1[pos1++];
            }else{
                ret[k++] = nums2[pos2++];
            }
        }
        if(pos1 == length1){
            while(pos2 < length2){
                ret[k++] = nums2[pos2++];
            }
        }
        if(pos2 == length2){
            while(pos1 < length1){
                ret[k++] = nums1[pos1++];
            }
        }
        return ((double)(ret[(length1+length2-1)/2]+ret[(length1+length2)/2]))/2;
    }

    //Q200:岛屿的个数
    //给定一个由 '1'（陆地）和 '0'（水）组成的的二维网格，计算岛屿的数量。一个岛被水包围，并且它是通过水平方向或垂直方向上相邻的陆地连接而成的。你可以假设网格的四个边均被水包围。
    //
    //示例 1:
    //
    //输入:
    //11110
    //11010
    //11000
    //00000
    //
    //输出: 1
    //示例 2:
    //
    //输入:
    //11000
    //11000
    //00100
    //00011
    //
    //输出: 3
    public int numIslands(char[][] grid) {
        int n = grid.length;
        if(n == 0) return 0;
        int m = grid[0].length;
        int count = 0;
        for(int i=0;i<n;i++){
            for(int j=0;j<m;j++){
                if(grid[i][j] == '1'){
                    dfs(i,j,grid);
                    count ++ ;
                }
            }
        }
        return count;
    }
    public void dfs(int i,int j,char[][] grid){
        grid[i][j] = '2';
        if(i > 0 && grid[i-1][j] == '1'){
            dfs(i-1,j,grid);
        }
        if(j > 0 && grid[i][j-1] == '1'){
            dfs(i,j-1,grid);
        }
        if(i < grid.length-1 && grid[i+1][j] == '1'){
            dfs(i+1,j,grid);
        }
        if(j < grid[0].length-1 && grid[i][j+1] == '1'){
            dfs(i,j+1,grid);
        }
    }

    //Q162:寻找峰值
    //峰值元素是指其值大于左右相邻值的元素。
    //
    //给定一个输入数组 nums，其中 nums[i] ≠ nums[i+1]，找到峰值元素并返回其索引。
    //
    //数组可能包含多个峰值，在这种情况下，返回任何一个峰值所在位置即可。
    //
    //你可以假设 nums[-1] = nums[n] = -∞。
    //
    //示例 1:
    //
    //输入: nums = [1,2,3,1]
    //输出: 2
    //解释: 3 是峰值元素，你的函数应该返回其索引 2。
    //示例 2:
    //
    //输入: nums = [1,2,1,3,5,6,4]
    //输出: 1 或 5
    //解释: 你的函数可以返回索引 1，其峰值元素为 2；
    //     或者返回索引 5， 其峰值元素为 6。
    //说明:
    //
    //你的解法应该是 O(logN) 时间复杂度的。
    public int findPeakElement(int[] nums) {
        if(nums.length == 1)
            return 0;
        if(nums[0] > nums[1])
            return 0;
        if(nums[nums.length - 1] > nums[nums.length - 2])
            return nums.length - 1;
        return binary(nums, 1 , nums.length - 2);
    }
    public int binary(int nums[], int left, int right){
        if(left > right)
            return -1;
        int mid = (left + right) / 2;
        if(nums[mid] > nums[mid - 1] && nums[mid] > nums[mid + 1])
            return mid;
        if(nums[mid - 1] > nums[mid + 1]){
            return binary(nums, left, mid - 1);
        }else{
            return binary(nums, mid + 1, right);
        }
    }

    //Q84:柱状图中最大的矩形
    //给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。
    //
    //求在该柱状图中，能够勾勒出来的矩形的最大面积。
    //以上是柱状图的示例，其中每个柱子的宽度为 1，给定的高度为 [2,1,5,6,2,3]。
    //图中阴影部分为所能勾勒出的最大矩形面积，其面积为 10 个单位。
    //示例:
    //
    //输入: [2,1,5,6,2,3]
    //输出: 10
    public int largestRectangleArea(int[] heights) {
        int length = heights.length;
        int[] heightsCopy = Arrays.copyOf(heights, length + 1);
        heightsCopy[length] = 0;
        Stack<Integer> stack  = new Stack<>();
        int maxArea = 0;
        for(int i=0;i<length+1;i++){
            while(!stack.empty() && heightsCopy[i]<heightsCopy[stack.peek()]){
                int top = stack.peek();
                stack.pop();
                maxArea = Math.max(maxArea, heightsCopy[top]*(stack.empty()?i:(i-stack.peek()-1)));
            }
            stack.push(i);
        }
        return maxArea;
    }

    //Q85:最大矩形
    //给定一个仅包含 0 和 1 的二维二进制矩阵，找出只包含 1 的最大矩形，并返回其面积。
    //
    //示例:
    //
    //输入:
    //[
    //  ["1","0","1","0","0"],
    //  ["1","0","1","1","1"],
    //  ["1","1","1","1","1"],
    //  ["1","0","0","1","0"]
    //]
    //输出: 6
    public int maximalRectangle(char[][] matrix) {
        return 0;
    }

    //Q5025:最长等差数列
    //给定一个整数数组 A，返回 A 中最长等差子序列的长度。
    //
    //回想一下，A 的子序列是列表 A[i_1], A[i_2], ..., A[i_k] 其中 0 <= i_1 < i_2 < ... < i_k <= A.length - 1。并且如果 B[i+1] - B[i]( 0 <= i < B.length - 1) 的值都相同，那么序列 B 是等差的。
    //
    //
    //
    //示例 1：
    //
    //输入：[3,6,9,12]
    //输出：4
    //解释：
    //整个数组是公差为 3 的等差数列。
    //示例 2：
    //
    //输入：[9,4,7,2,10]
    //输出：3
    //解释：
    //最长的等差子序列是 [4,7,10]。
    //示例 3：
    //
    //输入：[20,1,15,3,10,5,8]
    //输出：4
    //解释：
    //最长的等差子序列是 [20,15,10,5]。
    //
    //
    //提示：
    //
    //2 <= A.length <= 2000
    //0 <= A[i] <= 10000
    public int longestArithSeqLength(int[] A) {
        int n[][] = new int[A.length][20001];
        int ans = Integer.MIN_VALUE;
        int d;
        for(int i=0;i<A.length;i++){
            for(int j=0;j<i;j++){
                d = A[i]-A[j]+10000;
                if(n[j][d]==0){
                    n[i][d] = 2;
                }
                else {
                    n[i][d] = n[j][d]+1;
                }
                ans = Math.max(n[i][d],ans);
            }
        }
        return ans;
    }
}
