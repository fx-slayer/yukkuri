package test

import (
	"log"
	"testing"
)

func TestColor(t *testing.T){
	i :=  65420
	r := (i >> 8) & 0x00ff
	g := (i >> 4) & 0x00ff
	b := i & 0x00ff
	log.Println(r ,g ,b)
}

func TestSlice(t *testing.T){
	var s [][]int
	for i:=0 ;i<10 ;i++{
		si := make([]int ,0)
		for j:=0 ;j<5 ;j++{
			si = append(si, j)
		}
		log.Println(si)
		s = append(s, si)
	}

	//log.Println(s)
}
