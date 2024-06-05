package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
	"sync"
)

// 分词函数，处理大小写和标点
func tokenize(text string) []string {
	// 使用正则表达式移除标点符号
	reg := regexp.MustCompile(`[^\w\s]`)
	cleanText := reg.ReplaceAllString(text, "")
	// 将文本转为小写并分词
	words := strings.Fields(strings.ToLower(cleanText))
	return words
}

// 词频统计函数
func wordFrequency(text string) map[string]int {
	words := tokenize(text)
	frequency := make(map[string]int)
	for _, word := range words {
		frequency[word]++
	}
	return frequency
}

// 使用并发优化分词过程
func concurrentWordFrequency(text string) map[string]int {
	words := tokenize(text)
	frequency := make(map[string]int)
	var mu sync.Mutex
	var wg sync.WaitGroup

	for _, word := range words {
		wg.Add(1)
		go func(w string) {
			defer wg.Done()
			mu.Lock()
			frequency[w]++
			mu.Unlock()
		}(word)
	}

	wg.Wait()
	return frequency
}

// 词频排序函数
func sortWordFrequency(frequency map[string]int) []string {
	type kv struct {
		Key   string
		Value int
	}

	var sorted []kv
	for k, v := range frequency {
		sorted = append(sorted, kv{k, v})
	}

	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Value > sorted[j].Value
	})

	var result []string
	for _, kv := range sorted {
		result = append(result, fmt.Sprintf("%s: %d", kv.Key, kv.Value))
	}

	return result
}

func main() {
	// 从标准输入读取文本
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("请输入你的text:")
	text, _ := reader.ReadString('\n')

	// 词频统计
	frequency := concurrentWordFrequency(text)

	// 词频排序
	sortedFrequency := sortWordFrequency(frequency)

	// 输出结果
	fmt.Println("单词频率为:")
	for _, entry := range sortedFrequency {
		fmt.Println(entry)
	}
}
