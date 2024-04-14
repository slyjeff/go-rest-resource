package resource

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_EmbedMustEmbedChildResourceInParent(t *testing.T) {
	//arrange
	child := NewResource("child")
	parent := NewResource("parent")

	//act
	parent.EmbedResource("child", child)

	//assert
	a := assert.New(t)
	c, ok := parent.Embedded["child"]
	a.True(ok)
	a.Equal(child, c)
}

func Test_EmbedMustEmbedListOfChildResourcesInParent(t *testing.T) {
	//arrange
	child1 := NewResource("child")
	child2 := NewResource("child")
	children := []Resource{child1, child2}
	parent := NewResource("parent")

	//act
	parent.EmbedResources("children", children)

	//assert
	a := assert.New(t)
	c, ok := parent.Embedded["children"]
	a.True(ok)
	a.Equal(children, c)
}
