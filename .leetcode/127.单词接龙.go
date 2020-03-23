/*
 * @lc app=leetcode.cn id=127 lang=golang
 *
 * [127] 单词接龙
 *
 * https://leetcode-cn.com/problems/word-ladder/description/
 *
 * algorithms
 * Medium (35.55%)
 * Likes:    205
 * Dislikes: 0
 * Total Accepted:    21.1K
 * Total Submissions: 53.3K
 * Testcase Example:  '"hit"\n"cog"\n["hot","dot","dog","lot","log","cog"]'
 *
 * 给定两个单词（beginWord 和 endWord）和一个字典，找到从 beginWord 到 endWord
 * 的最短转换序列的长度。转换需遵循如下规则：
 *
 *
 * 每次转换只能改变一个字母。
 * 转换过程中的中间单词必须是字典中的单词。
 *
 *
 * 说明:
 *
 *
 * 如果不存在这样的转换序列，返回 0。
 * 所有单词具有相同的长度。
 * 所有单词只由小写字母组成。
 * 字典中不存在重复的单词。
 * 你可以假设 beginWord 和 endWord 是非空的，且二者不相同。
 *
 *
 * 示例 1:
 *
 * 输入:
 * beginWord = "hit",
 * endWord = "cog",
 * wordList = ["hot","dot","dog","lot","log","cog"]
 *
 * 输出: 5
 *
 * 解释: 一个最短转换序列是 "hit" -> "hot" -> "dot" -> "dog" -> "cog",
 * ⁠    返回它的长度 5。
 *
 *
 * 示例 2:
 *
 * 输入:
 * beginWord = "hit"
 * endWord = "cog"
 * wordList = ["hot","dot","dog","lot","log"]
 *
 * 输出: 0
 *
 * 解释: endWord "cog" 不在字典中，所以无法进行转换。
 *
 */
import "strings"

// @lc code=start
func getKey(chars []byte, wordLen, index int) string {
	var sb strings.Builder
	for j := 0; j < wordLen; j++ {
		if j == index {
			sb.WriteString("*")
		} else {
			sb.WriteByte(chars[j])
		}
	}
	return sb.String()
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	var wordLen, listLen, endWordInList = len(beginWord), len(wordList), false
	var preMap = make(map[string][]string, wordLen*listLen)
	for idx := 0; idx < listLen; idx++ {
		if !endWordInList && endWord == wordList[idx] {
			endWordInList = true
		}
		var chars = []byte(wordList[idx])
		for i := 0; i < wordLen; i++ {
			var key = getKey(chars, wordLen, i)
			if _, ok := preMap[key]; ok {
				preMap[key] = append(preMap[key], wordList[idx])
			} else {
				preMap[key] = []string{wordList[idx]}
			}
		}
	}
	if !endWordInList {
		return 0
	}

	//bfs
	visitedMap, queue, count := make(map[string]bool), []string{beginWord}, 0
	for len(queue) > 0 {
		count++
		newQueue := make([]string, 0)
		for _, item := range queue {
			for i := 0; i < wordLen; i++ {
				key := getKey([]byte(item), wordLen, i)
				for _, word := range preMap[key] {
					if word == endWord {
						return count + 1
					}
					if !visitedMap[word] {
						visitedMap[word] = true
						newQueue = append(newQueue, word)
					}
				}
			}
		}
		queue = newQueue
	}
	return 0
}

// @lc code=end

