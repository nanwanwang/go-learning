package calculate

import "testing"

func TestEven(t *testing.T) {
	//t.Log("Start TestEven")
	//t.Error("This is a T.error function.")
	if !Even(10) {
		t.Log("10 must be event")
		t.Fail()
	}

	if Even(77) {
		t.Log("77 is not even")
		t.Fail()
	}
}

func TestAdd(t *testing.T) {

	inputs := []struct{ a, b, c int }{
		{a: 1, b: 2, c: 3},
		{a: 4, b: 5, c: 9},
		{a: 10, b: 20, c: 30},
		{a: 3001, b: 4002, c: 7003},
	}
	for _, data := range inputs {
		if result := Add(data.a, data.b); result != data.c {
			t.Errorf("Add(%d,%d) expected result=%d,actual %d", data.a, data.b, data.c, result)
		}
	}
}

func TestOdd(t *testing.T) {
	if !Odd(11) {
		t.Log("11 must be odd")
		t.Fail()
	}

	if Odd(12) {
		t.Log("12 is not odd")
		t.Fail()
	}
}

func BenchmarkAdd(b *testing.B) {
	input := struct{ a, b, c int }{
		3001, 4002, 7003,
	}
	for i := 0; i < b.N; i++ {
		if result := Add(input.a, input.b); result != input.c {
			b.Errorf("Add(%d,%d) expected result=%d,actual result=%d\n",
				input.a, input.b, input.c, result)
		}
	}
}
