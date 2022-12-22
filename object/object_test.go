package object

import "testing"

func TestStringHashKeys(t *testing.T) {
  hello1 := &String{Value: "let's go pens"}
  hello2 := &String{Value: "let's go pens"}
  diff1 := &String{Value: "here we go steelers"}
  diff2 := &String{Value: "here we go steelers"}

  if hello1.Value.HashKey() != hello2.HashKey() {
    t.Errorf("strings with same content have different hash keys")
  }

  if diff1.Value.HashKey() != diff2.HashKey() {
    t.Errorf("strings with same content have different hash keys")
  }

  if hello1.Value.HashKey() == diff1.HashKey() {
    t.Errorf("strings with different content have same hash keys")
  }
}

