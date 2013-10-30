package cache

import (
	"testing"
)

var c *ConcurrentCache

type person struct {
	name     string
	age      int
	starsign string
}

var fred, barney person

func init() {

	c = New()
	fred = person{name: "Fred", age: 43, starsign: "Capricorn"}
	barney = person{name: "Barney", age: 52, starsign: "Pisces"}

}

func TestAdd(t *testing.T) {

	c.Add("fred", fred)
	c.Add("barney", barney)

	v, _ := c.Get("fred").(person) // type assertion from interface{} to person
	if v.name != "Fred" {
		t.Logf("Expected 'Fred', got %s", v)
		t.Fail()
	}

	v, _ = c.Get("barney").(person)
	if v.name != "Barney" {
		t.Logf("Expected 'Barney', got %s", v)
		t.Fail()
	}
}

func TestDelete(t *testing.T) {

	c.Delete("Fred")

	v := c.Get("Fred") // expecting nil which can't be asserted to type person
	if v != nil {
		t.Logf("Expected nil, got %s", v)
		t.Fail()
	}

	w, _ := c.Get("barney").(person)
	if w.name != "Barney" {
		t.Logf("Expected 'Barney', got %s", v)
		t.Fail()
	}

}

// func TestConcurrency(t *testing.T) {

// 	cc := New()

// 	go func{
// 		cc.Add("Fred", fred)
// 		time.Sleep(5)
// 		}()
// }
