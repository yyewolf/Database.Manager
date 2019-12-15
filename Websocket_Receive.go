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
			<td class="cell100"><p>Éleve</p></td>
			<td class="cell100"><p>Responsable (légal)</p></td>
			<td class="cell100"><p>Téléphone</p></td>
			<td class="cell100"><p>École</p></td>
			<td class="cell100"><p>Instrument</p></td>
			<td class="cell100"><p>Niveau Instrument</p></td>
			<td class="cell100"><p>Solfege</p></td>
			<td class="cell100"><p>Niveau Solfege</p></td>
			<td class="cell100"><p>Location</p></td>
			<td class="cell100"><p>Harmonie</p></td>
			<td class="cell100"><p>Mail</p></td>
			<td class="cell100"><p>Adresse</p></td>
			<td class="cell100"><p>Supprimer</p></td>
		</tr>
	</thead>`
	for rows.Next() {
		var r Column_Type
		rows.Scan(&r.ID, &r.Eleve, &r.Responsable, &r.Telephone, &r.Ecole, &r.Instrument, &r.Niveauinstrument, &r.Solfege, &r.Niveausolfege, &r.Location, &r.Harmonie, &r.Email, &r.Adresse)

		html += `<tr class="row100 body">
		<td class="cell100"><input id="todisable" onChange="Change('` + strconv.Itoa(r.ID) + `','eleve',this.value)" type="text" value="` + r.Eleve + `"/></td>
		<td class="cell100"><input id="todisable" onChange="Change('` + strconv.Itoa(r.ID) + `','responsable',this.value)" type="text" value="` + r.Responsable + `"/></td>
		<td class="cell100"><input id="todisable" onChange="Change('` + strconv.Itoa(r.ID) + `','telephone',this.value)" type="text" value="` + r.Telephone + `"/></td>
		<td class="cell100"><input id="todisable" onChange="Change('` + strconv.Itoa(r.ID) + `','ecole',this.value)" type="text" value="` + r.Ecole + `"/></td>
		<td class="cell100"><input id="todisable" onChange="Change('` + strconv.Itoa(r.ID) + `','instrument',this.value)" type="text" value="` + r.Instrument + `"/></td>
		<td class="cell100"><input id="todisable" onChange="Change('` + strconv.Itoa(r.ID) + `','niveauinstrument',this.value)" type="text" value="` + r.Niveauinstrument + `"/></td>
		<td class="cell100"><input id="todisable" onChange="Change('` + strconv.Itoa(r.ID) + `','solfege',this.value)" type="text" value="` + r.Solfege + `"/></td>
		<td class="cell100"><input id="todisable" onChange="Change('` + strconv.Itoa(r.ID) + `','niveausolfege',this.value)" type="text" value="` + r.Niveausolfege + `"/></td>
		<td class="cell100"><input id="todisable" onChange="Change('` + strconv.Itoa(r.ID) + `','location',this.value)" type="text" value="` + r.Location + `"/></td>
		<td class="cell100"><input id="todisable" onChange="Change('` + strconv.Itoa(r.ID) + `','harmonie',this.value)" type="text" value="` + r.Harmonie + `"/></td>
		<td class="cell100"><input id="todisable" onChange="Change('` + strconv.Itoa(r.ID) + `','email',this.value)" type="text" value="` + r.Email + `"/></td>
		<td class="cell100"><input id="todisable" onChange="Change('` + strconv.Itoa(r.ID) + `','adresse',this.value)" type="text" value="` + r.Adresse + `"/></td>
		<td class="cell100"><p> <button type="button" onClick="deleteentry('` + strconv.Itoa(r.ID) + `')" class="alert button">Retirer</button> </p></td>
		</tr>`
	}
	html += `</table>`
	//Send the HTTP message
	SendDB(html)
}

//THESE FUNCTIONS CONCERN THE ADD PAGE
//Filter the DB and sends it
func FilterDB_TabPage(Request Receive_Request) {
	rows, err := database.Query(Request.Request)
	if err != nil {
		return
	}
	var html string
	html = `<table class="hover" id="myTable">
	<thead>
		<tr class="row100 body thead">
			<td class="cell100"><p>Éleve</p></td>
			<td class="cell100"><p>Responsable (légal)</p></td>
			<td class="cell100"><p>Téléphone</p></td>
			<td class="cell100"><p>École</p></td>
			<td class="cell100"><p>Instrument</p></td>
			<td class="cell100"><p>Niveau Instrument</p></td>
			<td class="cell100"><p>Solfege</p></td>
			<td class="cell100"><p>Niveau Solfege</p></td>
			<td class="cell100"><p>Location</p></td>
			<td class="cell100"><p>Harmonie</p></td>
			<td class="cell100"><p>Mail</p></td>
			<td class="cell100"><p>Adresse</p></td>
			<td class="cell100"><p>Supprimer</p></td>
		</tr>
	</thead>`
	for rows.Next() {
		var r Column_Type
		rows.Scan(&r.ID, &r.Eleve, &r.Responsable, &r.Telephone, &r.Ecole, &r.Instrument, &r.Niveauinstrument, &r.Solfege, &r.Niveausolfege, &r.Location, &r.Harmonie, &r.Email, &r.Adresse)

		html += `<tr class="row100 body">
		<td class="cell100"><p>` + r.Eleve + `</p></td>
		<td class="cell100"><p>` + r.Responsable + `</p></td>
		<td class="cell100"><p>` + r.Telephone + `</p></td>
		<td class="cell100"><p>` + r.Ecole + `</p></td>
		<td class="cell100"><p>` + r.Instrument + `</p></td>
		<td class="cell100"><p>` + r.Niveauinstrument + `</p></td>
		<td class="cell100"><p>` + r.Solfege + `</p></td>
		<td class="cell100"><p>` + r.Niveausolfege + `</p></td>
		<td class="cell100"><p>` + r.Location + `</p></td>
		<td class="cell100"><p>` + r.Harmonie + `</p></td>
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
	statement, err := database.Prepare(`INSERT INTO DB (
		eleve,
		responsable,
		telephone,
		ecole,
		instrument,
		niveauinstrument,
		solfege,
		niveausolfege,
		location,
		harmonie,
		email,
		adresse
	)
	VALUES (
		"` + Request.Add.Eleve + `",
		"` + Request.Add.Responsable + `",
		"` + Request.Add.Telephone + `",
		"` + Request.Add.Ecole + `",
		"` + Request.Add.Instrument + `",
		"` + Request.Add.Niveauinstrument + `",
		"` + Request.Add.Solfege + `",
		"` + Request.Add.Niveausolfege + `",
		"` + Request.Add.Location + `",
		"` + Request.Add.Harmonie + `",
		"` + Request.Add.Email + `",
		"` + Request.Add.Adresse + `"
		);`)
	if err != nil {
		return
	}
	statement.Exec()
	//Send the DB over
	
	html := `<tr class="row100 body">
		<td class="cell100"><input id="todisable" onChange="Change('` + strconv.Itoa(Request.Add.ID) + `','eleve',this.value)" type="text" value="` + Request.Add.Eleve + `"/></td>
		<td class="cell100"><input id="todisable" onChange="Change('` + strconv.Itoa(Request.Add.ID) + `','responsable',this.value)" type="text" value="` + Request.Add.Responsable + `"/></td>
		<td class="cell100"><input id="todisable" onChange="Change('` + strconv.Itoa(Request.Add.ID) + `','telephone',this.value)" type="text" value="` + Request.Add.Telephone + `"/></td>
		<td class="cell100"><input id="todisable" onChange="Change('` + strconv.Itoa(Request.Add.ID) + `','ecole',this.value)" type="text" value="` + Request.Add.Ecole + `"/></td>
		<td class="cell100"><input id="todisable" onChange="Change('` + strconv.Itoa(Request.Add.ID) + `','instrument',this.value)" type="text" value="` + Request.Add.Instrument + `"/></td>
		<td class="cell100"><input id="todisable" onChange="Change('` + strconv.Itoa(Request.Add.ID) + `','niveauinstrument',this.value)" type="text" value="` + Request.Add.Niveauinstrument + `"/></td>
		<td class="cell100"><input id="todisable" onChange="Change('` + strconv.Itoa(Request.Add.ID) + `','solfege',this.value)" type="text" value="` + Request.Add.Solfege + `"/></td>
		<td class="cell100"><input id="todisable" onChange="Change('` + strconv.Itoa(Request.Add.ID) + `','niveausolfege',this.value)" type="text" value="` + Request.Add.Niveausolfege + `"/></td>
		<td class="cell100"><input id="todisable" onChange="Change('` + strconv.Itoa(Request.Add.ID) + `','location',this.value)" type="text" value="` + Request.Add.Location + `"/></td>
		<td class="cell100"><input id="todisable" onChange="Change('` + strconv.Itoa(Request.Add.ID) + `','harmonie',this.value)" type="text" value="` + Request.Add.Harmonie + `"/></td>
		<td class="cell100"><input id="todisable" onChange="Change('` + strconv.Itoa(Request.Add.ID) + `','email',this.value)" type="text" value="` + Request.Add.Email + `"/></td>
		<td class="cell100"><input id="todisable" onChange="Change('` + strconv.Itoa(Request.Add.ID) + `','adresse',this.value)" type="text" value="` + Request.Add.Adresse + `"/></td>
		<td class="cell100"><p> <button type="button" onClick="deleteentry('` + strconv.Itoa(Request.Add.ID) + `')" class="alert button">Retirer</button> </p></td>
		</tr>`
	SendSmallDB(html)
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
	Websocket_Receive_Functions["fdbtap"] = FilterDB_TabPage
	Websocket_Receive_Functions["deleteentry"] = DeleteEntry_AddPage
	Websocket_Receive_Functions["addentry"] = AddEntry_AddPage
	Websocket_Receive_Functions["change"] = EditEntry_AddPage
	Websocket_Receive_Functions["close"] = Close
}
