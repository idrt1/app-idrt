package main

import (
	"context"
	"database/sql"
	_ "modernc.org/sqlite"
	"myapp/client"
)

// App struct
type App struct {
	ctx           context.Context
	db            *sql.DB
	clientManager *client.ClientMananger
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup initializes the database
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx

	db, err := sql.Open("sqlite", "file:client.db?cache=shared")
	if err != nil {
		panic(err)
	}
	a.db = db

	a.clientManager.InitTable(a.db)

	//// Replace with the actual path to the Excel file
	//filePath := "./data.xlsx"
	//err = a.Insert(filePath)
	//if err != nil {
	//	fmt.Fprintf(os.Stderr, "Failed to insert data from Excel: %v\n", err)
	//} else {
	//	fmt.Println("Data insertion complete")
	//}
}

// shutdown closes the database connection
func (a *App) shutdown(ctx context.Context) {
	a.db.Close()
}

//func (a *App) Insert(filePath string) error {
//	if filePath == "" {
//		return fmt.Errorf("file path is empty")
//	}
//	// Check if the database is initialized
//	if a.db == nil {
//		return fmt.Errorf("database is not initialized")
//	}

//	// Open the Excel file
//	f, err := excelize.OpenFile(filePath)
//	if err != nil {
//		return err
//	}
//	defer f.Close()
//
//	// Get all the rows from the first sheet
//	rows, err := f.GetRows(f.GetSheetName(0))
//	if err != nil {
//		return err
//	}
//
//	// Iterate through rows and insert data into the database
//	for i, row := range rows {
//		// Skip header row
//		if i == 0 {
//			continue
//		}
//
//		// Assuming the columns in the Excel file match the following order:
//		// Nom, Prenom, AdresseElectronique, NumeroPort, etc.
//		if len(row) < 45 {
//			continue // Skip if there are not enough columns
//		}
//
//		_, err = a.db.Exec(`
//INSERT INTO client (Nom, Prenom, AdresseElectronique, NumeroPort, Categorie, ConjointSynd, CotiSpeciale, Cotisation, Cout, D_ou_N, DateCreation, DatePaiement, IDSyndique, NumeroTelProf, PremierAnCoti, Syndique, Titre, TypeInstallation, CategorieCotiCat, CategorieCotiDep, AdresseProf1, AdresseProf2, Age, Association, CodePostalProf, DateDiplome, DateModification, DateNaissance, DiplomeFaculte, LieuExercice, NumeroTelDomicile, PaysProf, PersAideAssistante, PersAssistante, PersCollaborateur, PersFemmeDeMenage, PersLaboratoire, PersReceptionniste, Personnels, Remarques, Responsable, Sexe, TypeExercice, VilleProf)
//VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
//			row[0], // Nom
//			row[1], // Prenom
//			row[2], // AdresseElectronique
//			row[3], // NumeroPort
//			row[4], // Categorie
//			row[5], // ConjointSynd
//			row[6], // CotiSpeciale
//			row[7], // Cotisation
//			row[8], // Cout
//			row[9], // D_ou_N
//			//row[10], // DateInstallation
//			row[11], // DateCreation
//			row[12], // DatePaiement
//			row[13], // IDSyndique
//			row[14], // NumeroTelProf
//			row[15], // PremierAnCoti
//			row[16], // Syndique
//			row[17], // Titre
//			row[18], // TypeInstallation
//			row[19], // CategorieCotiCat
//			row[20], // CategorieCotiDep
//			row[21], // AdresseProf1
//			row[22], // AdresseProf2
//			row[23], // Age
//			row[24], // Association
//			row[25], // CodePostalProf
//			row[26], // DateDiplome
//			row[27], // DateModification
//			row[28], // DateNaissance
//			row[29], // DiplomeFaculte
//			row[30], // LieuExercice
//			row[31], // NumeroTelDomicile
//			row[32], // PaysProf
//			row[33], // PersAideAssistante
//			row[34], // PersAssistante
//			row[35], // PersCollaborateur
//			row[36], // PersFemmeDeMenage
//			row[37], // PersLaboratoire
//			row[38], // PersReceptionniste
//			row[39], // Personnels
//			row[40], // Remarques
//			row[41], // Responsable
//			row[42], // Sexe
//			row[43], // TypeExercice
//			row[44], // VilleProf
//		)
//		if err != nil {
//			log.Printf("Error inserting row %d: %v", i, err)
//			continue
//		}
//	}
//
//	return nil
//}
