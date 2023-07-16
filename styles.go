package main

// Définit les styles utilisés par Client et Monitor

import (
	"fmt"
	lg "github.com/charmbracelet/lipgloss"
)

var (
	// barre d'état
	statusStyle = lg.NewStyle().
			Padding(0, 1, 0, 1).
			Margin(0, 1, 0, 1).
			Foreground(lg.Color("0")).
			Background(lg.Color("10"))
	// historique
	histStyle = lg.NewStyle().
			Padding(0, 1, 0, 1).
			Margin(0, 1, 0, 1).
			Foreground(lg.Color("10")).
			Background(lg.Color("0"))

	// affichage de la dernière commande
	// outputStyle = lg.NewStyle().
	// 		Padding(0, 1, 0, 1).
	// 		Margin(0, 1, 0, 1)

	// invite de commande
	inputStyle = lg.NewStyle().
			Padding(0, 1, 0, 1).
			Margin(0, 1, 0, 1)

	// fenêtre modale
	// modalStyle = lg.NewStyle().
	// 		Padding(0, 1, 0, 1).
	// 		Margin(0, 1, 0, 1).
	// 		BorderStyle(lg.DoubleBorder()).
	// 		BorderForeground(lg.Color("10"))

	focusFieldStyle = lg.NewStyle().
			BorderStyle(lg.NormalBorder()).
			BorderForeground(lg.Color("10"))

	unfocusFieldStyle = lg.NewStyle().
				BorderStyle(lg.NormalBorder()).
				BorderForeground(lg.Color("8"))

	// texte discret
	mutedTextStyle = lg.NewStyle().Foreground(lg.Color("8"))

	// texte normal
	normalTextStyle = lg.NewStyle().Foreground(lg.Color("15"))

	// curseur
	cursorStyle = lg.NewStyle().Reverse(true)

	// texte erreur
	errorTextStyle = lg.NewStyle().Foreground(lg.Color("9")).Padding(1)

	greenTextStyle  = lg.NewStyle().Foreground(lg.Color("0")).Background(lg.Color("10"))
	yellowTextStyle = lg.NewStyle().Foreground(lg.Color("0")).Background(lg.Color("11"))
	redTextStyle    = lg.NewStyle().Foreground(lg.Color("0")).Background(lg.Color("9"))
	invertTextStyle = lg.NewStyle().Reverse(true)
)

// Formate info connexion sur une ligne
var (
	conHeader = "ID: ____user/___login@server_________ w DNI [alert] LastCmd_______________________"
	// valeurs par défaut
	defUser    = "U--"
	defLogin   = "L--"
	defServer  = "---"
	defDNINo   = ""
	defDNIYes  = "DNI"
	defAlert   = " [--:--]"
	defLastCmd = "---"

	// formatter Login@Serveur
	cLenUser   = 8
	cLenLogin  = 8
	cLenServer = 15
	sLogServer = ": %8s/%8s@%-15s"

	// formatter Privilege/DNI
	sPrivDNI = " w %d/%3s"
	sDNI     = " w %3s"

	// formatter Alerte
	sAlert = " [%02d:%02d]"

	// formatter LastCmd
	cLenLastCmd = 30
	sLastCmd    = " %-30s"
)

func fmtConnexion(c *Console) string {
	conn := ""
	conn += fmt.Sprintf("%02d", c.ID)
	if c.IsConnected() {
		user := c.Session.User.Login
		if len(user) > cLenUser {
			user = user[:cLenUser-2] + ".."
		}
		login := c.Identity.Login
		if len(login) > cLenLogin {
			login = login[:cLenLogin-2] + ".."
		}

		server := c.Server.Address
		if len(server) > cLenServer {
			server = server[:cLenServer-3] + "..."
		}
		conn += fmt.Sprintf(sLogServer, user, login, server)

	} else {
		conn += fmt.Sprintf(sLogServer, defUser, defLogin, defServer)
	}

	if c.DNI {
		conn += fmt.Sprintf(sDNI /*c.Privilege,*/, defDNIYes)
	} else {
		conn += fmt.Sprintf(sDNI /*c.Privilege,*/, defDNINo)
	}
	if c.Alert {
		conn += fmt.Sprintf(sAlert,
			int(c.Countdown.Minutes()), int(c.Countdown.Seconds()))
	} else {
		conn += defAlert
	}

	if c.Alert {
		conn = redTextStyle.Render(conn)
	}

	lastCmd := defLastCmd
	if len(c.Results) > 0 {
		lastCmd = c.Results[len(c.Results)-1].Prompt
	}
	if len(lastCmd) > cLenLastCmd {
		lastCmd = lastCmd[:cLenLastCmd-3] + "..."
	}
	conn += fmt.Sprintf(sLastCmd, lastCmd)

	return conn
}

// Formate info registers sur une ligne
var (
	regHeader = "ID__ Serveur___ Description________________________ __________Value"

	// formatter
	rLenID     = 4
	rLenServer = 10
	rLenDesc   = 35
	rLenValue  = 15

	sReg = "%4d %-10s %-35s %15s"
)

func fmtRegister(r Register) string {

	server := r.Server
	if len(server) > rLenServer {
		server = server[:rLenServer-2] + ".."
	}
	desc := r.Description
	if len(desc) > rLenDesc {
		desc = desc[:rLenDesc-2] + ".."
	}
	value := r.State
	if len(value) > rLenValue {
		value = value[:rLenValue-2] + ".."
	}

	reg := fmt.Sprintf(sReg, r.ID, server, desc, value)
	return reg
}
