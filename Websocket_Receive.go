package main

import (
	"fmt"
	"os"
	"strconv"
)

//THESE FUNCTIONS CONCERN THE ADD PAGE
//Filter the DB and sends it
func FilterDB_AddPage(Request Receive_Request) {
	rows, err := database.Query(Request.Request)
	if err != nil {
		return
	}
	var html string
	html = `<table class="hover" id="myTable">
	<thead>
		<tr class="row100 body thead">
			<td class="cell100"><p>Nom</p></td>
			<td class="cell100"><p>Prénom</p></td>
			<td class="cell100"><p>Téléphone</p></td>
			<td class="cell100"><p>Instrument</p></td>
			<td class="cell100"><p>Niveau Instrument</p></td>
			<td class="cell100"><p>Solfege</p></td>
			<td class="cell100"><p>Niveau Solfege</p></td>
			<td class="cell100"><p>Mail</p></td>
			<td class="cell100"><p>Adresse</p></td>
			<td class="cell100"><p>Supprimer</p></td>
		</tr>
	</thead>`
	for rows.Next() {
		var r Column_Type
		rows.Scan(&r.ID, &r.Nom, &r.Prenom, &r.Telephone, &r.Instrument, &r.Niveauinstrument, &r.Solfege, &r.Niveausolfege, &r.Email, &r.Adresse)
		var solfege string
		if r.Solfege {
			solfege = "Oui"
		} else {
			solfege = "Non"
		}
		html += `<tr class="row100 body">
		<td class="cell100"><input id="todisable" onChange="Change('` + strconv.Itoa(r.ID) + `','nom',this.value)" type="text" value="` + r.Nom + `"</div></td>
		<td class="cell100"><input id="todisable" onChange="Change('` + strconv.Itoa(r.ID) + `','prenom',this.value)" type="text" value="` + r.Prenom + `"/></td>
		<td class="cell100"><input id="todisable" onChange="Change('` + strconv.Itoa(r.ID) + `','telephone',this.value)" type="text" value="` + r.Telephone + `"/></td>
		<td class="cell100"><input id="todisable" onChange="Change('` + strconv.Itoa(r.ID) + `','instrument',this.value)" type="text" value="` + r.Instrument + `"/></td>
		<td class="cell100"><input id="todisable" onChange="Change('` + strconv.Itoa(r.ID) + `','niveauinstrument',this.value)" type="text" value="` + r.Niveauinstrument + `"/></td>
		<td class="cell100"><input id="todisable" onChange="Change('` + strconv.Itoa(r.ID) + `','solfege',this.value)" type="text" value="` + solfege + `"/></td>
		<td class="cell100"><input id="todisable" onChange="Change('` + strconv.Itoa(r.ID) + `','niveausolfege',this.value)" type="text" value="` + r.Niveausolfege + `"/></td>
		<td class="cell100"><input id="todisable" onChange="Change('` + strconv.Itoa(r.ID) + `','email',this.value)" type="text" value="` + r.Email + `"/></td>
		<td class="cell100"><input id="todisable" onChange="Change('` + strconv.Itoa(r.ID) + `','adresse',this.value)" type="text" value="` + r.Adresse + `"/></td>
		<td class="cell100"><p> <button type="button" onClick="deleteentry('` + strconv.Itoa(r.ID) + `')" class="alert button">Retirer</button> </p></td>
		</tr>`
	}
	html += `</table>`
	//Send the HTTP message
	SendDB(html)
}

//Delete an entry from the DB
func DeleteEntry_AddPage(Request Receive_Request) {
	statement, err := database.Prepare(`DELETE FROM DB
		WHERE id=` + Request.ID + `;`)
	if err != nil {
		fmt.Println(err)
		return
	}
	statement.Exec()
	//Send the DB over
	FilterDB_AddPage(Request)
}

//Add an entry from the DB
func AddEntry_AddPage(Request Receive_Request) {
	var solfege string
	if Request.Add.SolfegeCheck == "Oui" {
		solfege = "1"
	} else {
		solfege = "0"
	}
	statement, err := database.Prepare(`INSERT INTO DB (nom, prenom, telephone, instrument, niveauinstrument, solfege, niveausolfege, email, adresse)
	VALUES (
		"` + Request.Add.Nom + `",
		"` + Request.Add.Prenom + `",
		"` + Request.Add.Telephone + `",
		"` + Request.Add.Instrument + `",
		"` + Request.Add.Niveauinstrument + `",
		"` + solfege + `",
		"` + Request.Add.Niveausolfege + `",
		"` + Request.Add.Email + `",
		"` + Request.Add.Adresse + `"
		);`)
	if err != nil {
		return
	}
	statement.Exec()
	//Send the DB over
	FilterDB_AddPage(Request)
}

//Add an entry from the DB
func EditEntry_AddPage(Request Receive_Request) {
	if Request.NewData == "Oui" {
		Request.NewData = "1"
	} else if Request.NewData == "Non" {
		Request.NewData = "0"
	}
	statement, err := database.Prepare(`UPDATE DB SET ` + Request.Column + `="` + Request.NewData + `" WHERE id=` + Request.ID + `;`)
	if err != nil {
		fmt.Println(err)
		return
	}
	statement.Exec()
}

//Closes the program
func Close(Request Receive_Request) {
	os.Exit(0)
}

func Websocket_Receive_AllFunctions() {
	// Directly translated from the old code.
	Websocket_Receive_Functions["fdbap"] = FilterDB_AddPage
	Websocket_Receive_Functions["deleteentry"] = DeleteEntry_AddPage
	Websocket_Receive_Functions["addentry"] = AddEntry_AddPage
	Websocket_Receive_Functions["change"] = EditEntry_AddPage
	Websocket_Receive_Functions["close"] = Close
}
