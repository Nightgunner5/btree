package btree

import "fmt"

type PhoneNumber struct {
	Number uint64
	Name   string
}

func (pn *PhoneNumber) Equals(other Sortable) bool {
	o, ok := other.(*PhoneNumber)
	return ok && pn.Number == o.Number
}

func (pn *PhoneNumber) LessThan(other Sortable) bool {
	o, ok := other.(*PhoneNumber)
	return ok && pn.Number < o.Number
}

func (pn *PhoneNumber) String() string {
	return fmt.Sprintf("%d %d-%d (%s)", pn.Number/10000000, (pn.Number/10000)%1000, pn.Number%10000, pn.Name)
}

func NewPhoneNumber(num uint64, name string) *PhoneNumber {
	if num < 10000000 {
		num += 5550000000
	}
	if num > 10000000000 {
		num = num % 10000000000
	}

	return &PhoneNumber{num, name}
}

func ExampleBTree_Find_phonebook() {
	phoneTree := new(BTree)

	phoneTree.Add(NewPhoneNumber(5551234, "Brett Bretterson"))
	phoneTree.Add(NewPhoneNumber(5555752, "So and So"))
	phoneTree.Add(NewPhoneNumber(15555551234, "Brett Bretterson <3"))
	phoneTree.Add(NewPhoneNumber(5555241, "Whats H. Face"))

	fmt.Println("Listing:")
	for it := phoneTree.Iter(); it.Valid(); it.Next() {
		fmt.Println(it.Value())
	}

	fmt.Println("Looking up 555-1234:")
	fmt.Println(phoneTree.Find(NewPhoneNumber(5551234, "")))

	fmt.Println("Looking up 555-1235:")
	fmt.Println(phoneTree.Find(NewPhoneNumber(5551235, "")))

	// Output:
	// Listing:
	// 555 555-1234 (Brett Bretterson)
	// 555 555-5241 (Whats H. Face)
	// 555 555-5752 (So and So)
	// Looking up 555-1234:
	// 555 555-1234 (Brett Bretterson)
	// Looking up 555-1235:
	// <nil>
}
