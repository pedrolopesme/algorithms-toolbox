package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSolutionWithRegularString(test *testing.T) {
	assert.Equal(test, []byte("ordeP"), Solution([]byte("Pedro")))
}

func TestSolutionWithEmptyString(test *testing.T) {
	assert.Equal(test, []byte(""), Solution([]byte("")))
}

func TestSolutionWithBigTest(test *testing.T) {
	assert.Equal(test,
		[]byte(".lsin mutnemref nasmucca ,tege anru lev ereusop ,siruam isin cenoD .siprut sumixam des sullesahP .ucra siuq tse etna dnefiele a ,lsin subinif alugil im ,tauqesnoc mutcid ni mes ,sumixam euqsetnelleP .isilicaf alluN .allignirf tema tis tse euqsetnellep sittam tU .taptulov euqitsirt sitrobol euqsiuQ .muspi reprocmallu ,eativ siprut arreviv ,muiterp neipas siuq maN .tile xe ue siuD .eranro rotittrop a susir ereusop tU .mes tege reprocmallu ,lev rutetcesnoc ni mutnemele ,sutem lsin cenoD .tile gnicsipida rutetcesnoc ,tema tis rolod muspi meroL"),
		Solution([]byte("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec nisl metus, elementum in consectetur vel, ullamcorper eget sem. Ut posuere risus a porttitor ornare. Duis eu ex elit. Nam quis sapien pretium, viverra turpis vitae, ullamcorper ipsum. Quisque lobortis tristique volutpat. Ut mattis pellentesque est sit amet fringilla. Nulla facilisi. Pellentesque maximus, sem in dictum consequat, mi ligula finibus nisl, a eleifend ante est quis arcu. Phasellus sed maximus turpis. Donec nisi mauris, posuere vel urna eget, accumsan fermentum nisl.")))
}