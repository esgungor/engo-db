package avl

import "testing"

func Test_AVLInsertLeftRotate(t *testing.T) {
	/*		-1
				b
			   / \
			  a   d
			     / \
			    c   e

	*/
	testRecord1 := Record{
		Key:   "a",
		Value: "test",
	}
	testRecord2 := Record{
		Key:   "b",
		Value: "test",
	}
	testRecord3 := Record{
		Key:   "c",
		Value: "test",
	}
	testRecord4 := Record{
		Key:   "d",
		Value: "test",
	}
	testRecord5 := Record{
		Key:   "e",
		Value: "test",
	}

	root := NewNode(&testRecord1)
	root = Insert(root, testRecord2)
	root = Insert(root, testRecord3)
	root = Insert(root, testRecord4)
	root = Insert(root, testRecord5)

	if root.record.Key != testRecord2.Key {
		t.Errorf("root value %v", root.record.Key)
		t.Fail()
	}
}

func Test_AVLInsertRightRotate(t *testing.T) {
	/*
			d     c
		   /     / \
		  c     b   d
		 /
		b

	*/
	testRecord1 := Record{
		Key:   "d",
		Value: "test",
	}
	testRecord2 := Record{
		Key:   "c",
		Value: "test",
	}
	testRecord3 := Record{
		Key:   "b",
		Value: "test",
	}

	root := NewNode(&testRecord1)
	root = Insert(root, testRecord2)
	root = Insert(root, testRecord3)

	if root.record.Key != testRecord2.Key {
		t.Errorf("root value %v", root.record.Key)
		t.Fail()
	}
}
