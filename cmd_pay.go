package main

import (
	"fmt"
	"strconv"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type BalanceMsg struct{}

type PayMsg struct {
	To     string
	Amount int
}

var bank = Cmd{
	Name:      "bank",
	ShortHelp: "effectuer des opérations bancaires",
	SubCmds: []Cmd{
		{
			Path:      []string{"bank"},
			Name:      "balance",
			ShortHelp: "affiche le solde du compte",
			Parse: func(args []string) any {
				return BalanceMsg{}
			},
		},
		{
			Path:      []string{"bank"},
			Name:      "pay",
			ShortHelp: "effectue un transfert de monnaie",
			Args: []Arg{
				{Name: "to", ShortHelp: "compte à créditer"},
				{Name: "amount", ShortHelp: "montant à transférer"},
			},
			Parse: func(args []string) any {
				to := args[0]
				if amount, err := strconv.Atoi(args[1]); err != nil {
					return Eval{
						Error: fmt.Errorf("Amount : %w", errInvalidArgument),
					}
				} else {
					return PayMsg{to, amount}
				}

			},
		},
	},
}

type PayModel struct {
	To       textinput.Model
	Amount   textinput.Model
	Password textinput.Model
}

func (p *PayModel) Init() tea.Cmd {
	return nil
}

func (p *PayModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return p, nil
}

func (p *PayModel) View() string {
	return ""
}
