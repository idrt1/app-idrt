package client

import (
	"database/sql"
)

type ClientMananger struct {
	DB *sql.DB
}

type Client struct {
	ID                  int    `json:"id"`
	Nom                 string `json:"nom"`
	Prenom              string `json:"prenom"`
	AdresseElectronique string `json:"adresseElectronique"`
	NumeroPort          string `json:"numeroPort"`
	Categorie           string `json:"categorie"`
	ConjointSynd        string `json:"conjointSynd"`
	CotiSpeciale        string `json:"cotiSpeciale"`
	Cotisation          string `json:"cotisation"`
	Cout                string `json:"cout"`
	D_ou_N              string `json:"d_ou_N"`
	DateInstallation    string `json:"dateInstallation"`
	DateCreation        string `json:"dateCreation"`
	DatePaiement        string `json:"datePaiement"`
	IDSyndique          string `json:"idSyndique"`
	NumeroTelProf       string `json:"numeroTelProf"`
	PremierAnCoti       string `json:"premierAnCoti"`
	Syndique            string `json:"syndique"`
	Titre               string `json:"titre"`
	TypeInstallation    string `json:"typeInstallation"`
	CategorieCotiCat    string `json:"categorieCotiCat"`
	CategorieCotiDep    string `json:"categorieCotiDep"`
	AdresseProf1        string `json:"adresseProf1"`
	AdresseProf2        string `json:"adresseProf2"`
	Age                 string `json:"age"`
	Association         string `json:"association"`
	CodePostalProf      string `json:"codePostalProf"`
	DateDiplome         string `json:"dateDiplome"`
	DateModification    string `json:"dateModification"`
	DateNaissance       string `json:"dateNaissance"`
	DiplomeFaculte      string `json:"diplomeFaculte"`
	LieuExercice        string `json:"lieuExercice"`
	NumeroTelDomicile   string `json:"numeroTelDomicile"`
	PaysProf            string `json:"paysProf"`
	PersAideAssistante  string `json:"persAideAssistante"`
	PersAssistante      string `json:"persAssistante"`
	PersCollaborateur   string `json:"persCollaborateur"`
	PersFemmeDeMenage   string `json:"persFemmeDeMenage"`
	PersLaboratoire     string `json:"persLaboratoire"`
	PersReceptionniste  string `json:"persReceptionniste"`
	Personnels          string `json:"personnels"`
	Remarques           string `json:"remarques"`
	Responsable         string `json:"responsable"`
	Sexe                string `json:"sexe"`
	TypeExercice        string `json:"typeExercice"`
	VilleProf           string `json:"villeProf"`
}

func (cm *ClientMananger) InitTable(db *sql.DB) {
	cm.DB = db
	_, err := cm.DB.Exec(`
CREATE TABLE IF NOT EXISTS client (
    id INTEGER PRIMARY KEY, 
    Nom TEXT, 
    Prenom TEXT, 
    AdresseElectronique TEXT, 
    NumeroPort TEXT, 
    Categorie TEXT, 
    ConjointSynd TEXT, 
    CotiSpeciale TEXT, 
    Cotisation TEXT, 
    Cout TEXT, 
    D_ou_N TEXT, 
    DateInstallation DATE, 
    DateCreation DATE, 
    DatePaiement DATE, 
    IDSyndique INTEGER, 
    NumeroTelProf TEXT, 
    PremierAnCoti INTEGER, 
    Syndique TEXT, 
    Titre TEXT, 
    TypeInstallation TEXT, 
    CategorieCotiCat TEXT, 
    CategorieCotiDep TEXT, 
    AdresseProf1 TEXT, 
    AdresseProf2 TEXT, 
    Age INTEGER, 
    Association TEXT, 
    CodePostalProf TEXT, 
    DateDiplome DATE, 
    DateModification DATE, 
    DateNaissance DATE, 
    DiplomeFaculte TEXT, 
    LieuExercice TEXT, 
    NumeroTelDomicile TEXT, 
    PaysProf TEXT, 
    PersAideAssistante TEXT, 
    PersAssistante TEXT, 
    PersCollaborateur TEXT, 
    PersFemmeDeMenage TEXT, 
    PersLaboratoire TEXT, 
    PersReceptionniste TEXT, 
    Personnels TEXT, 
    Remarques TEXT, 
    Responsable TEXT, 
    Sexe TEXT, 
    TypeExercice TEXT, 
    VilleProf TEXT
)`)

	if err != nil {
		panic(err)
	}
}

func (cm *ClientMananger) InsertClient(client Client) error {
	_, err := cm.DB.Exec(`
INSERT INTO client (Nom, Prenom, AdresseElectronique, NumeroPort, Categorie, ConjointSynd, CotiSpeciale, Cotisation, Cout, D_ou_N, DateCreation, DatePaiement, IDSyndique, NumeroTelProf, PremierAnCoti, Syndique, Titre, TypeInstallation, CategorieCotiCat, CategorieCotiDep, AdresseProf1, AdresseProf2, Age, Association, CodePostalProf, DateDiplome, DateModification, DateNaissance, DiplomeFaculte, LieuExercice, NumeroTelDomicile, PaysProf, PersAideAssistante, PersAssistante, PersCollaborateur, PersFemmeDeMenage, PersLaboratoire, PersReceptionniste, Personnels, Remarques, Responsable, Sexe, TypeExercice, VilleProf)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		client.Nom,                 // Nom
		client.Prenom,              // Prenom
		client.AdresseElectronique, // AdresseElectronique
		client.NumeroPort,          // NumeroPort
		client.Categorie,           // Categorie
		client.ConjointSynd,        // ConjointSynd
		client.CotiSpeciale,        // CotiSpeciale
		client.Cotisation,          // Cotisation
		client.Cout,                // Cout
		client.D_ou_N,              // D_ou_N
		//client.DateInstallation, // DateInstallation
		client.DateCreation,       // DateCreation
		client.DatePaiement,       // DatePaiement
		client.IDSyndique,         // IDSyndique
		client.NumeroTelProf,      // NumeroTelProf
		client.PremierAnCoti,      // PremierAnCoti
		client.Syndique,           // Syndique
		client.Titre,              // Titre
		client.TypeInstallation,   // TypeInstallation
		client.CategorieCotiCat,   // CategorieCotiCat
		client.CategorieCotiDep,   // CategorieCotiDep
		client.AdresseProf1,       // AdresseProf1
		client.AdresseProf2,       // AdresseProf2
		client.Age,                // Age
		client.Association,        // Association
		client.CodePostalProf,     // CodePostalProf
		client.DateDiplome,        // DateDiplome
		client.DateModification,   // DateModification
		client.DateNaissance,      // DateNaissance
		client.DiplomeFaculte,     // DiplomeFaculte
		client.LieuExercice,       // LieuExercice
		client.NumeroTelDomicile,  // NumeroTelDomicile
		client.PaysProf,           // PaysProf
		client.PersAideAssistante, // PersAideAssistante
		client.PersAssistante,     // PersAssistante
		client.PersCollaborateur,  // PersCollaborateur
		client.PersFemmeDeMenage,  // PersFemmeDeMenage
		client.PersLaboratoire,    // PersLaboratoire
		client.PersReceptionniste, // PersReceptionniste
		client.Personnels,         // Personnels
		client.Remarques,          // Remarques
		client.Responsable,        // Responsable
		client.Sexe,               // Sexe
		client.TypeExercice,       // TypeExercice
		client.VilleProf,          // VilleProf
	)
	if err != nil {
		return err
	}
	return nil
}

func (cm *ClientMananger) GetAllClient() ([]Client, error) {
	rows, err := cm.DB.Query(`SELECT id, Nom, Prenom, Categorie, AdresseElectronique FROM client`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var clients []Client
	for rows.Next() {
		var client Client
		err := rows.Scan(
			&client.ID,                  // id
			&client.Nom,                 // Nom
			&client.Prenom,              // Prenom
			&client.Categorie,           // Categorie
			&client.AdresseElectronique, // AdresseElectronique
		)
		if err != nil {
			return nil, err
		}
		clients = append(clients, client)
	}
	return clients, nil
}

func (cm *ClientMananger) GetClientByID(id int) (Client, error) {
	row := cm.DB.QueryRow(`SELECT id, Nom, Prenom, AdresseElectronique, NumeroPort, Categorie, ConjointSynd, CotiSpeciale, Cotisation, Cout, D_ou_N, DateCreation, DatePaiement, IDSyndique, NumeroTelProf, PremierAnCoti, Syndique, Titre, TypeInstallation, CategorieCotiCat, CategorieCotiDep, AdresseProf1, AdresseProf2, Age, Association, CodePostalProf, DateDiplome, DateModification, DateNaissance, DiplomeFaculte, LieuExercice, NumeroTelDomicile, PaysProf, PersAideAssistante, PersAssistante, PersCollaborateur, PersFemmeDeMenage, PersLaboratoire, PersReceptionniste, Personnels, Remarques, Responsable, Sexe, TypeExercice, VilleProf FROM client WHERE id = ?`, id)

	var client Client
	err := row.Scan(
		&client.ID,                  // id
		&client.Nom,                 // Nom
		&client.Prenom,              // Prenom
		&client.AdresseElectronique, // AdresseElectronique
		&client.NumeroPort,          // NumeroPort
		&client.Categorie,           // Categorie
		&client.ConjointSynd,        // ConjointSynd
		&client.CotiSpeciale,        // CotiSpeciale
		&client.Cotisation,          // Cotisation
		&client.Cout,                // Cout
		&client.D_ou_N,              // D_ou_N
		//&client.DateInstallation,    // DateInstallation
		&client.DateCreation,       // DateCreation
		&client.DatePaiement,       // DatePaiement
		&client.IDSyndique,         // IDSyndique
		&client.NumeroTelProf,      // NumeroTelProf
		&client.PremierAnCoti,      // PremierAnCoti
		&client.Syndique,           // Syndique
		&client.Titre,              // Titre
		&client.TypeInstallation,   // TypeInstallation
		&client.CategorieCotiCat,   // CategorieCotiCat
		&client.CategorieCotiDep,   // CategorieCotiDep
		&client.AdresseProf1,       // AdresseProf1
		&client.AdresseProf2,       // AdresseProf2
		&client.Age,                // Age
		&client.Association,        // Association
		&client.CodePostalProf,     // CodePostalProf
		&client.DateDiplome,        // DateDiplome
		&client.DateModification,   // DateModification
		&client.DateNaissance,      // DateNaissance
		&client.DiplomeFaculte,     // DiplomeFaculte
		&client.LieuExercice,       // LieuExercice
		&client.NumeroTelDomicile,  // NumeroTelDomicile
		&client.PaysProf,           // PaysProf
		&client.PersAideAssistante, // PersAideAssistante
		&client.PersAssistante,     // PersAssistante
		&client.PersCollaborateur,  // PersCollaborateur
		&client.PersFemmeDeMenage,  // PersFemmeDeMenage
		&client.PersLaboratoire,    // PersLaboratoire
		&client.PersReceptionniste, // PersReceptionniste
		&client.Personnels,         // Personnels
		&client.Remarques,          // Remarques
		&client.Responsable,        // Responsable
		&client.Sexe,               // Sexe
		&client.TypeExercice,       // TypeExercice
		&client.VilleProf,          // VilleProf
	)
	if err != nil {
		return Client{}, err
	}
	return client, nil
}
