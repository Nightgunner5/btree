// Binary Tree
package btree

import "fmt"

type Value_T int64

type BTree struct {
	Left   *BTree
	Value  Value_T
	Right  *BTree
	isset  bool
	length int
}

// Adds a value to the BTree while keeping the tree sorted.
// Each value is only added once -- duplicates are ignored.
func (tree *BTree) Add(val Value_T) bool {
	if tree.isset {
		tree.length = 0 // Reset lengths
		if tree.Value == val {
			return false
		}
		if tree.Value > val {
			if tree.Left == nil {
				tree.Left = new(BTree)
			}
			return tree.Left.Add(val)
		}
		if tree.Right == nil {
			tree.Right = new(BTree)
		}
		return tree.Right.Add(val)
	}
	tree.isset = true
	tree.Value = val
	return true

}

// Special case only used when both Left and Right of the parent are set but Value isn't.
// This function is called on Right in all cases.
func (tree *BTree) removeFirst() Value_T {
	if tree.Left != nil {
		value := tree.Left.Iter().Value()
		tree.Left.Remove(value)
		return value
	}
	value := tree.Value
	tree.Remove(value)
	return value
}

// Removes a value from the BTree. Returns true if the BTree contained the value.
// This function will move values around in the BTree if a branch would lose its value while it still had children.
func (tree *BTree) Remove(val Value_T) bool {
	if !tree.isset {
		return false
	}
	tree.length = 0 // Reset lengths
	if tree.Value == val {
		if tree.Right != nil {
			if tree.Left == nil {
				tree.Value = tree.Right.Value
				tree.Left = tree.Right.Left
				tree.Right = tree.Right.Right
			} else {
				tree.Value = tree.Right.removeFirst()
			}
		} else if tree.Left != nil {
			tree.Value = tree.Left.Value
			tree.Right = tree.Left.Right
			tree.Left = tree.Left.Left
		} else {
			tree.isset = false
		}
		return true
	}
	if tree.Value > val {
		if tree.Left != nil {
			defer func() {
				if !tree.Left.isset {
					tree.Left = nil
				}
			}()
			return tree.Left.Remove(val)
		}
		return false
	}
	if tree.Right != nil {
		defer func() {
			if !tree.Right.isset {
				tree.Right = nil
			}
		}()
		return tree.Right.Remove(val)
	}
	return false
}

// Returns true if the value exists in the BTree.
func (tree *BTree) Contains(val Value_T) bool {
	if val == tree.Value {
		return true
	}
	if val < tree.Value && tree.Left != nil {
		return tree.Left.Contains(val)
	}
	if val > tree.Value && tree.Right != nil {
		return tree.Right.Contains(val)
	}
	return false
}

// Returns the number of elements in this BTree.
func (tree *BTree) Length() int {
	if tree.isset {
		if tree.length == 0 {
			tree.length = tree.uncachedLength()
		}
		return tree.length
	}
	return 0
}

func (tree *BTree) cachedLength() int {
	length := 0
	if tree.isset {
		length++
		if tree.Left != nil {
			length += tree.Left.Length()
		}
		if tree.Right != nil {
			length += tree.Right.Length()
		}
	}
	return length
}

func (tree *BTree) uncachedLength() int {
	length := 0
	if tree.isset {
		length++
		if tree.Left != nil {
			length += tree.Left.uncachedLength()
		}
		if tree.Right != nil {
			length += tree.Right.uncachedLength()
		}
	}
	return length
}

// Returns a new iterator located at the start of the BTree.
func (tree *BTree) Iter() *Iterator {
	if tree.Left != nil {
		return &Iterator{tree.Left.Iter(), tree, 0}
	}
	return &Iterator{nil, tree, 1}
}

// Copies the BTree into a slice using an Iterator.
func (tree *BTree) ToSlice() []Value_T {
	if !tree.isset {
		return nil
	}

	slice := make([]Value_T, tree.Length())
	i := 0
	for iter := tree.Iter(); iter.Valid(); iter.Next() {
		slice[i] = iter.Value()
		i++
	}
	return slice
}

// Returns a text representation of the result of the ToSlice() function.
func (tree *BTree) String() string {
	return fmt.Sprintf("%v", tree.ToSlice())
}

// Returns a balanced copy of the current BTree, where the maximum depth is as low as possible.
func (tree *BTree) Balance() *BTree {
	slice := tree.ToSlice()

	balanced := new(BTree)

	balanced._balance(slice)

	return balanced
}

func (tree *BTree) _balance(slice []Value_T) {
	if len(slice) == 0 {
		return
	}
	if len(slice) == 1 {
		tree.Add(slice[0])
		return
	}
	mid := len(slice) / 2
	tree.Add(slice[mid])
	tree._balance(slice[:mid])
	tree._balance(slice[mid:])
}

type Iterator struct {
	child *Iterator
	self  *BTree
	index byte
}

// Returns true if Value() can be called successfully.
// More specifically, it returns true if the iterator's current value is within the bounds of the BTree.
func (iter *Iterator) Valid() bool {
	return iter.index < 3
}

// Advances the iterator to the next value. This has no effect if the current value is invalid.
func (iter *Iterator) Next() {
	switch iter.index {
	case 0:
		iter.child.Next()
		if iter.child.Valid() {
			return
		}
		iter.index = 1
		iter.child = nil

	case 1:
		if iter.self.Right != nil {
			iter.child = iter.self.Right.Iter()
			iter.index = 2
		} else {
			iter.index = 3
		}

	case 2:
		iter.child.Next()
		if iter.child.Valid() {
			return
		}
		iter.index = 3
		iter.child = nil
	}
}

// Returns the current value represented by this iterator.
// Panics if Valid() would return false.
func (iter *Iterator) Value() Value_T {
	if !iter.Valid() {
		panic("Value() called on invalid Iterator")
	}
	if iter.index == 1 {
		return iter.self.Value
	}
	return iter.child.Value()
}
