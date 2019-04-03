package emojipedia

type Emojidex map[string]*Emoji

func (emojidex *Emojidex) Add(name string, emoji *Emoji) {
	(*emojidex)[name] = emoji
}
