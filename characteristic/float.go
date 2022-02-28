package characteristic

import (
	"github.com/brutella/hap/log"

	"net/http"
)

type Float struct {
	*C
}

func NewFloat(t string) *Float {
	c := New()
	c.Type = t
	return &Float{c}
}

// SetValue sets a value
func (c *Float) SetValue(v float64) {
	c.setValue(v, nil)
}

func (c *Float) SetMinValue(v float64) {
	c.MinVal = v
}

func (c *Float) SetMaxValue(v float64) {
	c.MaxVal = v
}

func (c *Float) SetStepValue(v float64) {
	c.StepVal = v
}

// Value returns the value of c as float64.
func (c *Float) Value() float64 {
	v, _ := c.C.valueRequest(nil)
	if v == nil {
		return 0
	}

	return v.(float64)
}

func (c *Float) MinValue() float64 {
	return c.MinVal.(float64)
}

func (c *Float) MaxValue() float64 {
	return c.MaxVal.(float64)
}

func (c *Float) StepValue() float64 {
	return c.StepVal.(float64)
}

// OnSetRemoteValue set c.SetValueRequestFunc and calls fn only
// if the value is going to be updated from a request.
func (c *Float) OnSetRemoteValue(fn func(v float64) error) {
	c.SetValueRequestFunc = func(v interface{}, r *http.Request) int {
		if r == nil {
			return 0
		}

		if err := fn(v.(float64)); err != nil {
			log.Debug.Println(err)
			return -70402
		}
		return 0
	}
}

// OnValueRemoteUpdate calls fn when the value of the characteristic was updated.
// If the provided http request is not nil, the value was updated by a client (ex. iOS device).
func (c *Float) OnValueUpdate(fn func(new, old float64, r *http.Request)) {
	c.OnCValueUpdate(func(c *C, new, old interface{}, r *http.Request) {
		fn(new.(float64), old.(float64), r)
	})
}

// OnValueRemoteUpdate calls fn when the value of the characteristic was updated by a client.
func (c *Float) OnValueRemoteUpdate(fn func(v float64)) {
	c.OnCValueUpdate(func(c *C, new, old interface{}, r *http.Request) {
		if r != nil {
			fn(new.(float64))
		}
	})
}
