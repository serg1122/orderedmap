package orderedmap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_OrderedMap(t *testing.T) {

	m1Key1 := `m1Key1`
	m1Value1 := `m1Value1`

	m1Key2 := `m1Key2`
	// m1Value2 := `m1Value2`

	m1 := CreateOrderedMap()

	assert.IsType(t, &OrderedMap{}, m1)
	assert.Equal(t, 0, m1.Length())

	assert.False(t, m1.HasKey(m1Key1))
	assert.False(t, m1.HasKey(m1Key2))

	v1FirstNodeNil, err1FirstNode := m1.FirstNode()

	assert.Nil(t, v1FirstNodeNil)
	assert.IsType(t, &ErrorChainIsEmpty{}, err1FirstNode)

	v1LastNodeNil, err1LastNode := m1.LastNode()

	assert.Nil(t, v1LastNodeNil)
	assert.IsType(t, &ErrorChainIsEmpty{}, err1LastNode)

	m1.Set(m1Key1, m1Value1)

	assert.Equal(t, 1, m1.Length())
	assert.True(t, m1.HasKey(m1Key1))
	assert.False(t, m1.HasKey(m1Key2))

	m1FirstNode, m1ErrFirstNil := m1.FirstNode()

	assert.IsType(t, &Node{}, m1FirstNode)
	assert.Nil(t, m1ErrFirstNil)
}
