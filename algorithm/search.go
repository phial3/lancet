// Copyright 2021 dudaodong@gmail.com. All rights reserved.
// Use of this source code is governed by MIT license

// Package algorithm contain some basic algorithm functions. eg. sort, search, list, linklist, stack, queue, tree, graph. TODO
package algorithm

import "github.com/duke-git/lancet/lancetconstraints"

// BinarySearch search for target within a sorted slice, recursive call itself.
// If a target is found, the index of the target is returned. Else the function return -1
func BinarySearch[T any](slice []T, target T, lowIndex, highIndex int, comparator lancetconstraints.Comparator) int {
	if highIndex < lowIndex || len(slice) == 0 {
		return -1
	}

	midIndex := int(lowIndex + (highIndex-lowIndex)/2)
	// slice[midIndex] > target
	if comparator.Compare(slice[midIndex], target) == 1 {
		return BinarySearch(slice, target, lowIndex, midIndex-1, comparator)
	} else if comparator.Compare(slice[midIndex], target) == -1 {
		return BinarySearch(slice, target, midIndex+1, highIndex, comparator)
	}

	return midIndex
}

// BinaryIterativeSearch search for target within a sorted slice.
// If a target is found, the index of the target is returned. Else the function return -1
func BinaryIterativeSearch[T any](slice []T, target T, lowIndex, highIndex int, comparator lancetconstraints.Comparator) int {
	startIndex := lowIndex
	endIndex := highIndex

	var midIndex int
	for startIndex <= endIndex {
		midIndex = int(startIndex + (endIndex-startIndex)/2)
		if comparator.Compare(slice[midIndex], target) == 1 {
			endIndex = midIndex - 1
		} else if comparator.Compare(slice[midIndex], target) == -1 {
			startIndex = midIndex + 1
		} else {
			return midIndex
		}
	}
	return -1
}
