// Binary Tree
package btree

import (
	"fmt"
	"reflect"
)

type Sortable interface {
	Equals(other Sortable) bool
	LessThan(other Sortable) bool
}

type BTree struct {
	left   *BTree
	value  value_t
	right  *BTree
	isset  bool
	length int
}

func wrap(val interface{}) value_t {
	// Okay, get ready for a long, annoying function caused by the lack of generics!

	if v, ok := val.(Sortable); ok {
		return &value_impl_Sortable{v}
	}
	if v, ok := val.(int); ok {
		return &value_impl_int{v}
	}
	if v, ok := val.(uint); ok {
		return &value_impl_uint{v}
	}
	if v, ok := val.(int8); ok {
		return &value_impl_int8{v}
	}
	if v, ok := val.(uint8); ok {
		return &value_impl_uint8{v}
	}
	if v, ok := val.(int16); ok {
		return &value_impl_int16{v}
	}
	if v, ok := val.(uint16); ok {
		return &value_impl_uint16{v}
	}
	if v, ok := val.(int32); ok {
		return &value_impl_int32{v}
	}
	if v, ok := val.(uint32); ok {
		return &value_impl_uint32{v}
	}
	if v, ok := val.(int64); ok {
		return &value_impl_int64{v}
	}
	if v, ok := val.(uint64); ok {
		return &value_impl_uint64{v}
	}
	panic(fmt.Sprintln("Unknown datatype ", reflect.TypeOf(val), " in ", val))
}

// Adds a value to the BTree while keeping the tree sorted.
// Each value is only added once -- duplicates are ignored.
func (tree *BTree) Add(val interface{}) bool {
	return tree.add(wrap(val))
}

func (tree *BTree) add(val value_t) bool {
	if tree.isset {
		tree.length = 0 // Reset lengths
		if tree.value.Equals(val) {
			return false
		}
		if val.LessThan(tree.value) {
			if tree.left == nil {
				tree.left = new(BTree)
			}
			return tree.left.add(val)
		}
		if tree.right == nil {
			tree.right = new(BTree)
		}
		return tree.right.add(val)
	}
	tree.isset = true
	tree.value = val
	return true

}

// Special case only used when both Left and Right of the parent are set but Value isn't.
func (tree *BTree) removeFirst() value_t {
	if tree.left != nil {
		value := tree.left.removeFirst()
		if !tree.left.isset {
			tree.left = nil
		}
		return value
	}
	value := tree.value
	tree.remove(value)
	return value
}

func (tree *BTree) replaceWith(other *BTree) {
	tree.value = other.value
	tree.left = other.left
	tree.right = other.right
}

// Removes a value from the BTree. Returns true if the BTree contained the value.
// This function will move values around in the BTree if a branch would lose its value while it still had children.
func (tree *BTree) Remove(val interface{}) bool {
	return tree.remove(wrap(val))
}

func (tree *BTree) remove(val value_t) bool {
	if !tree.isset {
		return false
	}
	tree.length = 0 // Reset lengths
	if tree.value.Equals(val) {
		if tree.right != nil {
			if tree.left == nil {
				tree.replaceWith(tree.right)
			} else {
				tree.value = tree.right.removeFirst()
			}
		} else if tree.left != nil {
			tree.replaceWith(tree.left)
		} else {
			tree.isset = false
		}
		return true
	}
	if val.LessThan(tree.value) {
		if tree.left != nil {
			defer func() {
				if !tree.left.isset {
					tree.left = nil
				}
			}()
			return tree.left.remove(val)
		}
		return false
	}
	if tree.right != nil {
		defer func() {
			if !tree.right.isset {
				tree.right = nil
			}
		}()
		return tree.right.remove(val)
	}
	return false
}

// Returns true if the value exists in the BTree.
func (tree *BTree) Contains(val interface{}) bool {
	return tree.contains(wrap(val))
}

func (tree *BTree) contains(val value_t) bool {
	if val.Equals(tree.value) {
		return true
	}
	if val.LessThan(tree.value) && tree.left != nil {
		return tree.left.contains(val)
	}
	if tree.value.LessThan(val) && tree.right != nil {
		return tree.right.contains(val)
	}
	return false
}

// Returns the value as it exists in the BTree or nil if it does not occur in the BTree.
func (tree *BTree) Find(val interface{}) interface{} {
	return tree.find(wrap(val))
}

func (tree *BTree) find(val value_t) interface{} {
	if val.Equals(tree.value) {
		return tree.value.Value()
	}
	if val.LessThan(tree.value) && tree.left != nil {
		return tree.left.find(val)
	}
	if tree.value.LessThan(val) && tree.right != nil {
		return tree.right.find(val)
	}
	return nil
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
		if tree.left != nil {
			length += tree.left.Length()
		}
		if tree.right != nil {
			length += tree.right.Length()
		}
	}
	return length
}

func (tree *BTree) uncachedLength() int {
	length := 0
	if tree.isset {
		length++
		if tree.left != nil {
			length += tree.left.uncachedLength()
		}
		if tree.right != nil {
			length += tree.right.uncachedLength()
		}
	}
	return length
}

// Returns a new iterator located at the start of the BTree.
func (tree *BTree) Iter() *Iterator {
	if tree.left != nil {
		return &Iterator{tree.left.Iter(), tree, 0}
	}
	return &Iterator{nil, tree, 1}
}

// Copies the BTree into a slice using an Iterator.
func (tree *BTree) ToSlice() []interface{} {
	if !tree.isset {
		return nil
	}

	slice := make([]interface{}, tree.Length())
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

func (tree *BTree) _balance(slice []interface{}) {
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
		if iter.self.right != nil {
			iter.child = iter.self.right.Iter()
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
func (iter *Iterator) Value() interface{} {
	if !iter.Valid() {
		panic("Value() called on invalid Iterator")
	}
	if iter.index == 1 {
		return iter.self.value.Value()
	}
	return iter.child.Value()
}
