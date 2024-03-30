package task_155

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinStack(t *testing.T) {
	ms := Constructor()
	ms.Push(-2)
	ms.Push(0)
	ms.Push(-3)
	assert.Equal(t, ms.GetMin(), -3)
	ms.Pop()
	assert.Equal(t, ms.Top(), 0)
	assert.Equal(t, ms.GetMin(), -2)
}
