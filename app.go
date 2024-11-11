package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"github.com/xuri/excelize/v2"
	"log"
	_ "modernc.org/sqlite"
	"myapp/client"
	"path/filepath"
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
}

// shutdown closes the database connection
func (a *App) shutdown(ctx context.Context) {
	a.db.Close()
}

func (a *App) Insert() error {
	filePath, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{})
	if err != nil {
		return err
	}
	if filePath == "" {
		return fmt.Errorf("file path is empty")
	}
	// Check if the database is initialized
	if a.db == nil {
		return fmt.Errorf("database is not initialized")
	}

	// Open the Excel file
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	// Get all the rows from the first sheet
	rows, err := f.GetRows(f.GetSheetName(0))
	if err != nil {
		return err
	}

	for i, row := range rows {
		// Skip header row
		if i == 0 {
			continue
		}

		if len(row) < 45 {
			continue
		}

		_, err = a.db.Exec(`
INSERT INTO client (Nom, Prenom, AdresseElectronique, NumeroPort, Categorie, ConjointSynd, CotiSpeciale, Cotisation, Cout, D_ou_N, DateCreation, DatePaiement, IDSyndique, NumeroTelProf, PremierAnCoti, Syndique, Titre, TypeInstallation, CategorieCotiCat, CategorieCotiDep, AdresseProf1, AdresseProf2, Age, Association, CodePostalProf, DateDiplome, DateModification, DateNaissance, DiplomeFaculte, LieuExercice, NumeroTelDomicile, PaysProf, PersAideAssistante, PersAssistante, PersCollaborateur, PersFemmeDeMenage, PersLaboratoire, PersReceptionniste, Personnels, Remarques, Responsable, Sexe, TypeExercice, VilleProf)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
			row[0], // Nom
			row[1], // Prenom
			row[2], // AdresseElectronique
			row[3], // NumeroPort
			row[4], // Categorie
			row[5], // ConjointSynd
			row[6], // CotiSpeciale
			row[7], // Cotisation
			row[8], // Cout
			row[9], // D_ou_N
			//row[10], // DateInstallation
			row[11], // DateCreation
			row[12], // DatePaiement
			row[13], // IDSyndique
			row[14], // NumeroTelProf
			row[15], // PremierAnCoti
			row[16], // Syndique
			row[17], // Titre
			row[18], // TypeInstallation
			row[19], // CategorieCotiCat
			row[20], // CategorieCotiDep
			row[21], // AdresseProf1
			row[22], // AdresseProf2
			row[23], // Age
			row[24], // Association
			row[25], // CodePostalProf
			row[26], // DateDiplome
			row[27], // DateModification
			row[28], // DateNaissance
			row[29], // DiplomeFaculte
			row[30], // LieuExercice
			row[31], // NumeroTelDomicile
			row[32], // PaysProf
			row[33], // PersAideAssistante
			row[34], // PersAssistante
			row[35], // PersCollaborateur
			row[36], // PersFemmeDeMenage
			row[37], // PersLaboratoire
			row[38], // PersReceptionniste
			row[39], // Personnels
			row[40], // Remarques
			row[41], // Responsable
			row[42], // Sexe
			row[43], // TypeExercice
			row[44], // VilleProf
		)
		if err != nil {
			log.Printf("Error inserting row %d: %v", i, err)
			continue
		}
	}
	return nil
}

func (a *App) Export() error {
	filePath, err := runtime.SaveFileDialog(a.ctx, runtime.SaveDialogOptions{
		Title: "Enregistrer le fichier",
		Filters: []runtime.FileFilter{
			{DisplayName: "Fichier Excel", Pattern: "*.xlsx"},
		},
	})
	if err != nil {
		return fmt.Errorf("Erreur lors de l'ouverture du dialogue de sauvegarde: %w", err)
	}
	if filePath == "" {
		return fmt.Errorf("Chemin de fichier non spécifié")
	}
	if filepath.Ext(filePath) != ".xlsx" {
		filePath += ".xlsx"
	}

	// Initialisation du fichier Excel
	f := excelize.NewFile()
	defer f.Close()

	// Liste des en-têtes de colonne
	headers := []string{
		"Nom", "Prenom", "AdresseElectronique", "NumeroPort", "Categorie",
		"ConjointSynd", "CotiSpeciale", "Cotisation", "Cout", "D_ou_N", "DateInstallation",
		"DateCreation", "DatePaiement", "IDSyndique", "NumeroTelProf",
		"PremierAnCoti", "Syndique", "Titre", "TypeInstallation", "CategorieCotiCat",
		"CategorieCotiDep", "AdresseProf1", "AdresseProf2", "Age", "Association",
		"CodePostalProf", "DateDiplome", "DateModification", "DateNaissance",
		"DiplomeFaculte", "LieuExercice", "NumeroTelDomicile", "PaysProf",
		"PersAideAssistante", "PersAssistante", "PersCollaborateur", "PersFemmeDeMenage",
		"PersLaboratoire", "PersReceptionniste", "Personnels", "Remarques", "Responsable",
		"Sexe", "TypeExercice", "VilleProf",
	}

	for i, header := range headers {
		cell := getColumnName(i) + "1"
		if err := f.SetCellValue("Sheet1", cell, header); err != nil {
			return fmt.Errorf("Erreur lors de l'insertion de l'en-tête %s à la cellule %s: %w", header, cell, err)
		}
	}

	rows, err := a.db.Query("SELECT Nom, Prenom, AdresseElectronique, NumeroPort, Categorie, ConjointSynd, CotiSpeciale, Cotisation, Cout, D_ou_N, DateInstallation, DateCreation, DatePaiement, IDSyndique, NumeroTelProf, PremierAnCoti, Syndique, Titre, TypeInstallation, CategorieCotiCat, CategorieCotiDep, AdresseProf1, AdresseProf2, Age, Association, CodePostalProf, DateDiplome, DateModification, DateNaissance, DiplomeFaculte, LieuExercice, NumeroTelDomicile, PaysProf, PersAideAssistante, PersAssistante, PersCollaborateur, PersFemmeDeMenage, PersLaboratoire, PersReceptionniste, Personnels, Remarques, Responsable, Sexe, TypeExercice, VilleProf FROM client")
	if err != nil {
		return fmt.Errorf("Erreur lors de la récupération des données de la base: %w", err)
	}
	defer rows.Close()

	rowIndex := 2
	for rows.Next() {
		var data [45]sql.NullString

		if err := rows.Scan(&data[0], &data[1], &data[2], &data[3], &data[4], &data[5], &data[6], &data[7], &data[8], &data[9],
			&data[10], &data[11], &data[12], &data[13], &data[14], &data[15], &data[16], &data[17], &data[18], &data[19],
			&data[20], &data[21], &data[22], &data[23], &data[24], &data[25], &data[26], &data[27], &data[28], &data[29],
			&data[30], &data[31], &data[32], &data[33], &data[34], &data[35], &data[36], &data[37], &data[38], &data[39],
			&data[40], &data[41], &data[42], &data[43], &data[44]); err != nil {
			log.Printf("Erreur lors de la lecture de la ligne %d: %v", rowIndex, err)
			return err
		}

		// Insérer les valeurs dans les colonnes respectives
		for colIndex, col := range data {
			cell := getColumnName(colIndex) + fmt.Sprintf("%d", rowIndex)
			value := ""
			if col.Valid {
				value = col.String // Si non-NULL, récupérer la valeur
			}
			if err := f.SetCellValue("Sheet1", cell, value); err != nil {
				log.Printf("Erreur lors de l'insertion de la valeur '%s' à la cellule %s: %v", value, cell, err)
				return err
			}
		}
		rowIndex++
	}

	// Sauvegarde du fichier
	if err := f.SaveAs(filePath); err != nil {
		return fmt.Errorf("Erreur lors de la sauvegarde du fichier : %w", err)
	}

	log.Println("Fichier exporté avec succès à l'emplacement :", filePath)
	return nil
}

func getColumnName(index int) string {
	result := ""
	for index >= 0 {
		remainder := index % 26
		result = string('A'+remainder) + result
		index = (index / 26) - 1
	}
	return result
}
