package main

type Model interface {
	Save()
	Load(id uint64)
}
