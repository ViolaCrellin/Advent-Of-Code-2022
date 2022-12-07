package solutions

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"example.com/adventofcode/util"
	"example.com/adventofcode/util/trees"
)

var all = []*trees.DoublyLinkedValueNode{}

func Seven(input string) string {
	consoleLog := strings.Split(input, "\n")
	root := &trees.DoublyLinkedValueNode{
		Parent:   nil,
		Id:       "root",
		Size:     0,
		Children: make(map[string]*trees.DoublyLinkedValueNode),
	}

	BuildFileStructure(consoleLog, root)
	unusedSpace := 70000000 - root.Size
	additionalSpaceNeeded := 30000000 - unusedSpace
	acceptableSizes := map[string]int{}
	candidatesForDeletion := map[string]int{}
	for i := range root.Children {
		FindDirectoriesSmallerThan(root.Children[i], acceptableSizes, 100000)
		FindDirectoriesLargerThan(root.Children[i], candidatesForDeletion, additionalSpaceNeeded)
	}

	bySize := util.MapValuesInt(candidatesForDeletion)
	sort.IntSlice(bySize).Sort()
	return fmt.Sprintf("Part 1: %d, Part 2: %d", util.SumMapValues(acceptableSizes, true), bySize[0])
}

func FindDirectoriesSmallerThan(node *trees.DoublyLinkedValueNode, acceptableSizes map[string]int, size int) {
	for i := range node.Children {
		if node.Size <= size {
			acceptableSizes[node.Path] = node.Size
		}

		FindDirectoriesSmallerThan(node.Children[i], acceptableSizes, size)
	}
}

func FindDirectoriesLargerThan(node *trees.DoublyLinkedValueNode, acceptableSizes map[string]int, size int) {
	for i := range node.Children {
		if node.Size >= size {
			acceptableSizes[node.Path] = node.Size
		}

		FindDirectoriesLargerThan(node.Children[i], acceptableSizes, size)
	}
}

func BuildFileStructure(consoleLog []string, node *trees.DoublyLinkedValueNode) bool {
	for i := 0; i <= len(consoleLog); i++ {
		line := consoleLog[i]
		// Iterate outputted file listing
		if line == "$ ls" {
			for {
				i++
				if i >= len(consoleLog) {
					return true
				}
				listLine := consoleLog[i]
				if strings.HasPrefix(listLine, "$") {
					break
				}

				if strings.HasPrefix(listLine, "dir") {
					childDir := strings.Replace(listLine, "dir ", "", 1)
					_, ok := node.Children[childDir]
					if !ok {
						dirChild := &trees.DoublyLinkedValueNode{
							Parent:   node,
							Id:       childDir,
							Path:     fmt.Sprintf("%s.%s", node.Path, childDir),
							Size:     0,
							Children: make(map[string]*trees.DoublyLinkedValueNode),
						}

						all = append(all, dirChild)
						node.Children[childDir] = dirChild
					}
				} else {
					r := regexp.MustCompile(`(?P<fileSize>\d+)\s(?P<fileName>.+)`)
					match := r.FindStringSubmatch(listLine)
					values := util.GetRegexMapOfNamedCaptureGroupValues(r, match)
					fileName := values["fileName"]
					fileSize, err := strconv.Atoi(values["fileSize"])
					if err != nil {
						//meh
					}
					// Doesnt matter if it already exists or not. We just overwrite it
					fileChild := &trees.DoublyLinkedValueNode{
						Parent:   node,
						Id:       fileName,
						Path:     fmt.Sprintf("%s.%s", node.Path, fileName),
						Size:     fileSize,
						Children: nil,
					}

					// traverse back up adding this thing's size to all parents
					CasadeSizeIncrease(fileChild, fileSize)
					all = append(all, fileChild)
					node.Children[fileName] = fileChild
				}
			}
		}

		if i >= len(consoleLog) {
			return true
		}

		line = consoleLog[i]
		if line == "$ cd .." {
			if node.Parent != nil {
				return BuildFileStructure(consoleLog[i+1:], node.Parent)
			}
		}

		r := regexp.MustCompile(`\$\scd\s[a-zA-Z]+`)
		if r.MatchString(line) {
			moveIntoDirId := strings.Replace(line, "$ cd ", "", 1)
			moveIntoDir, ok := node.Children[moveIntoDirId]
			if ok {
				return BuildFileStructure(consoleLog[i+1:], moveIntoDir)
			}
		}
	}

	return true
}

func CasadeSizeIncrease(node *trees.DoublyLinkedValueNode, size int) {
	if node == nil || node.Parent == nil {
		return
	}
	node.Parent.Size += size
	CasadeSizeIncrease(node.Parent, size)
}
