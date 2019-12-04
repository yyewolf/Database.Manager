package main

import (
	"fmt"
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
		<td class="cell100"><p>` + r.Nom + `</p></td>
		<td class="cell100"><p>` + r.Prenom + `</p></td>
		<td class="cell100"><p>` + r.Telephone + `</p></td>
		<td class="cell100"><p>` + r.Instrument + `</p></td>
		<td class="cell100"><p>` + r.Niveauinstrument + `</p></td>
		<td class="cell100"><p>` + solfege + `</p></td>
		<td class="cell100"><p>` + r.Niveausolfege + `</p></td>
		<td class="cell100"><p>` + r.Email + `</p></td>
		<td class="cell100"><p>` + r.Adresse + `</p></td>
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

func Websocket_Receive_AllFunctions() {
	// Directly translated from the old code.
	Websocket_Receive_Functions["fdbap"] = FilterDB_AddPage
	Websocket_Receive_Functions["deleteentry"] = DeleteEntry_AddPage
	Websocket_Receive_Functions["addentry"] = AddEntry_AddPage
}
