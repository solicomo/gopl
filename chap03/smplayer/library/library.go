package library

import "errors"

type Music struct {
	Id     string
	Name   string
	Artist string
	Source string
	Type   string
}

type MusicManager struct {
	musics []Music
}

func NewMusicManager() *MusicManager {
	return &MusicManager{make([]Music, 0)}
}

func (m *MusicManager) Len() int {
	return len(m.musics)
}

func (m *MusicManager) Get(index int) (music *Music, err error) {	
	if index < 0 || index >= m.Len() {
		err = errors.New("No such music")
		return
	}

	music = &m.musics[index]
	return
}

func (m *MusicManager) Find(name string) *Music {
	for _, music := range m.musics {
		if music.Name == name {
			return &music
		}
	}

	return nil
}

func (m *MusicManager) Add(music *Music) {
	m.musics = append(m.musics, *music)
}

func (m *MusicManager) Del(index int) *Music {
	if index < 0 || index >= m.Len() {
		return nil
	}

	music := &m.musics[index]

	m.musics = append(m.musics[:index], m.musics[index+1:])

	return music
}


