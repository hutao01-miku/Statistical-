package main

import (
	"reflect"
	"testing"
)

func TestTokenize(t *testing.T) {
	text := "Hello, world! This is a test."
	expected := []string{"hello", "world", "this", "is", "a", "test"}
	result := tokenize(text)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestWordFrequency(t *testing.T) {
	text := "Hello world hello"
	expected := map[string]int{"hello": 2, "world": 1}
	result := wordFrequency(text)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestConcurrentWordFrequency(t *testing.T) {
	text := "Hello world hello"
	expected := map[string]int{"hello": 2, "world": 1}
	result := concurrentWordFrequency(text)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestSortWordFrequency(t *testing.T) {
	frequency := map[string]int{"hello": 2, "world": 1}
	expected := []string{"hello: 2", "world: 1"}
	result := sortWordFrequency(frequency)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}
