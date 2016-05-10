package library

import (
	"errors"
	"fmt"
)

type MusicEntry struct {
	Id string

	Name string

	Artist string

	Source string

	Type string
}

type MusicManager struct {
	musics []MusicEntry
}

func NewMusicManager() *MusicManager {

	return &MusicManager{make([]MusicEntry, 0)}

}

func (m *MusicManager) Len() int {

	return len(m.musics)

}

func (m *MusicManager) Get(index int) (music *MusicEntry, err error) {

	if index < 0 || index >= len(m.musics) {

		return nil, errors.New("Index out of range.")

	}
	//fmt.Println(m)
	return &m.musics[index], nil

}

func (m *MusicManager) Find(name string) *MusicEntry {

	if len(m.musics) == 0 {

		return nil

	}

	for _, m := range m.musics {

		if m.Name == name {

			return &m

		}

	}

	return nil

}

func (m *MusicManager) Add(music *MusicEntry) {

	m.musics = append(m.musics, *music)

}

func (m *MusicManager) Remove(index int) *MusicEntry {

	if index < 0 || index > len(m.musics) {
		fmt.Println("请重新选择删除的序号..")
		return nil

	}

	removedMusic := &m.musics[index-1]

	// 从数组切片中删除元素

	if index < len(m.musics) { // 中间元素
		m.musics = append(m.musics[:index-1], m.musics[index:]...)
	} else { // 删除的是最后一个元素
		//fmt.Println("删除最后一个")
		m.musics = m.musics[:index-1]

	}

	return removedMusic

}
