package models

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/arnavsurve/gomo/pkg/styles"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"gopkg.in/yaml.v2"
)

var (
	home, _ = os.UserHomeDir()
	Path    = fmt.Sprintf("%s/.config/gomo/config.yaml", home)
)

type Config struct {
	Focus int
	Short int
	Long  int
}

func (c *Config) GetConf() *Config {
	file, err := os.ReadFile(Path)
	if err != nil {
		log.Fatalf("error reading config file at %s: %v", Path, err)
	}

	err = yaml.Unmarshal(file, c)
	if err != nil {
		log.Fatalf("unmarshal: %v", err)
	}

	return c
}

type configModel struct {
	focusIndex int
	inputs     []textinput.Model
}

func NewConfigModel() configModel {
	m := configModel{
		inputs: make([]textinput.Model, 3),
	}

	m.saveConfig(Path)
	config := Config{}
	c := config.GetConf()

	var t textinput.Model
	for i := range m.inputs {
		t = textinput.New()
		t.Cursor.Style = styles.CursorStyle
		t.CharLimit = 32

		switch i {
		case 0:
			t.Placeholder = strconv.Itoa(c.Focus)
			t.Focus()
			t.TextStyle = styles.FocusedStyle
		case 1:
			t.Placeholder = strconv.Itoa(c.Short)
		case 2:
			t.Placeholder = strconv.Itoa(c.Long)
		}

		m.inputs[i] = t
	}

	return m
}

func (m configModel) Init() tea.Cmd {
	return nil
}

func (m configModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc", "q":
			return m, tea.Quit

		// Set focus to next input
		case "tab", "shift+tab", "enter", "up", "down":
			s := msg.String()

			// Submit if enter pressed while last index (submit button) is selected
			if s == "enter" && m.focusIndex == len(m.inputs) {
				if err := m.saveConfig(Path); err != nil {
					log.Fatalf("Error saving config: %v", err)
				}
				return m, tea.Quit
			}

			// Cycle indexes
			if s == "up" || s == "shift+tab" {
				m.focusIndex--
			} else {
				m.focusIndex++
			}

			if m.focusIndex > len(m.inputs) {
				m.focusIndex = 0
			} else if m.focusIndex < 0 {
				m.focusIndex = len(m.inputs)
			}

			cmds := make([]tea.Cmd, len(m.inputs))
			for i := 0; i <= len(m.inputs)-1; i++ {
				if i == m.focusIndex {
					// Set focused state
					cmds[i] = m.inputs[i].Focus()
					m.inputs[i].PromptStyle = styles.FocusedStyle
					m.inputs[i].TextStyle = styles.FocusedStyle
					continue
				}
				// Remove focused state
				m.inputs[i].Blur()
				m.inputs[i].PromptStyle = styles.NoStyle
				m.inputs[i].TextStyle = styles.NoStyle
			}

			return m, tea.Batch(cmds...)
		}
	}

	// Handle character input
	cmd := m.updateInputs(msg)

	return m, cmd
}

func (m *configModel) updateInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(m.inputs))

	// Only text inputs with Focus() set will respond, so it's safe to simply
	// update all of them here without any further logic.
	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}

	return tea.Batch(cmds...)
}

// saveConfig writes the current input values to the config file
func (m configModel) saveConfig(path string) error {
	// Create directory if it doesn't exist
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return fmt.Errorf("error creating config directory: %w", err)
	}

	// Create or open the config file
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("error opening/creating config file: %w", err)
	}
	defer file.Close()

	// Encode config to YAML
	enc := yaml.NewEncoder(file)
	defer enc.Close()

	focus, err := strconv.Atoi(m.inputs[0].Value())
	if err != nil || m.inputs[0].Value() == "" {
		focus = 25
	}
	short, err := strconv.Atoi(m.inputs[1].Value())
	if err != nil || m.inputs[1].Value() == "" {
		short = 10
	}
	long, err := strconv.Atoi(m.inputs[2].Value())
	if err != nil || m.inputs[2].Value() == "" {
		long = 20
	}

	err = enc.Encode(Config{
		Focus: focus,
		Short: short,
		Long:  long,
	})
	if err != nil {
		return fmt.Errorf("error encoding yaml: %w", err)
	}

	return nil
}

func (m configModel) View() string {
	var b strings.Builder

	b.WriteRune('\n')
	b.WriteString("Focus " + m.inputs[0].View() + " min\n")
	b.WriteString("Short break " + m.inputs[1].View() + " min\n")
	b.WriteString("Long break " + m.inputs[2].View() + " min")

	button := &styles.BlurredButton
	if m.focusIndex == len(m.inputs) {
		button = &styles.FocusedButton
	}
	fmt.Fprintf(&b, "\n\n%s\n\n", *button)
	return b.String()
}
