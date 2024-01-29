package main

import tea "github.com/charmbracelet/bubbletea"

/* UPDATE */
func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	if !m.initialSelected {
		switch msg := msg.(type) {
		case tea.WindowSizeMsg:
			m.height = msg.Height
			m.width = msg.Width
		case tea.KeyMsg:
			switch msg.String() {
			case "ctrl+c":
				return m, tea.Quit
			case "up", "k":
				if m.cursor > 0 {
					m.cursor--
				}
			case "down", "j":
				if m.cursor < len(m.choices)-1 {
					m.cursor++
				}
			case "enter", "space":
				m.initialSelection = m.cursor
				m.initialSelected = true
				// Initialize or focus the text input here if needed
			}
		}
	} else if m.initialSelection == 0 && m.location.zipcode == "" {
		/* ZIPCODE INPUT */ 
		in := m.location.input
		var cmd tea.Cmd
		
		switch msg := msg.(type) {
		case tea.KeyMsg:
			ms := msg.String()
			val := in.Value()
			// Check if the key press is a digit before updating the input model
			if len(val) < 5 && isDigit(msg) || ms == "backspace" || ms == "ctrl+c" || ms == "enter" {
				in, cmd = in.Update(msg)
				m.location.input = in
			}
	
			if ms == "ctrl+c" {
				return m, tea.Quit
			}
	
			if len(val) == 5 && ms == "enter" || ms == " " {
				if len(in.Value()) == 5 {
					m.location.zipcode = in.Value()
					in.SetValue("") // Reset the input field after capturing the zipcode
				}
			}
		}
		return m, cmd

	}


	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.height = msg.Height
		m.width = msg.Width
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		}
	}

	return m, cmd
}
func isDigit(keyMsg tea.KeyMsg) bool {
    return keyMsg.Type == tea.KeyRunes && len(keyMsg.Runes) == 1 && keyMsg.Runes[0] >= '0' && keyMsg.Runes[0] <= '9'
}