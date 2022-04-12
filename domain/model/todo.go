package model

import "errors"

type Todo struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func NewTodo(title, content string) (*Todo, error) {
	if title == "" && content == "" {
		return nil, errors.New("titleとcontentを入力してください")
	} else if title == "" {
		return nil, errors.New("titleを入力してください")
	} else if content == "" {
		return nil, errors.New("contentを入力してください")
	}

	todo := &Todo{
		Title: title,
		Content: content,
	}

	return todo, nil
}

func (t *Todo) Set(title, content string) error {
	if title != "" && content != "" {
		t.Title = title
		t.Content = content
		return nil
	} else if title != "" {
		t.Title = title
		return nil
	} else if content != "" {
		t.Content = content
		return nil
	}

	return errors.New("更新する内容を入力してください")
}
