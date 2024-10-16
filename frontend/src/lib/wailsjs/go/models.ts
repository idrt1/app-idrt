export namespace client {
	
	export class Client {
	    id: number;
	    nom: string;
	    prenom: string;
	    adresseElectronique: string;
	    numeroPort: string;
	    categorie: string;
	    conjointSynd: string;
	    cotiSpeciale: string;
	    cotisation: string;
	    cout: string;
	    d_ou_N: string;
	    dateInstallation: string;
	    dateCreation: string;
	    datePaiement: string;
	    idSyndique: string;
	    numeroTelProf: string;
	    premierAnCoti: string;
	    syndique: string;
	    titre: string;
	    typeInstallation: string;
	    categorieCotiCat: string;
	    categorieCotiDep: string;
	    adresseProf1: string;
	    adresseProf2: string;
	    age: string;
	    association: string;
	    codePostalProf: string;
	    dateDiplome: string;
	    dateModification: string;
	    dateNaissance: string;
	    diplomeFaculte: string;
	    lieuExercice: string;
	    numeroTelDomicile: string;
	    paysProf: string;
	    persAideAssistante: string;
	    persAssistante: string;
	    persCollaborateur: string;
	    persFemmeDeMenage: string;
	    persLaboratoire: string;
	    persReceptionniste: string;
	    personnels: string;
	    remarques: string;
	    responsable: string;
	    sexe: string;
	    typeExercice: string;
	    villeProf: string;
	
	    static createFrom(source: any = {}) {
	        return new Client(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.nom = source["nom"];
	        this.prenom = source["prenom"];
	        this.adresseElectronique = source["adresseElectronique"];
	        this.numeroPort = source["numeroPort"];
	        this.categorie = source["categorie"];
	        this.conjointSynd = source["conjointSynd"];
	        this.cotiSpeciale = source["cotiSpeciale"];
	        this.cotisation = source["cotisation"];
	        this.cout = source["cout"];
	        this.d_ou_N = source["d_ou_N"];
	        this.dateInstallation = source["dateInstallation"];
	        this.dateCreation = source["dateCreation"];
	        this.datePaiement = source["datePaiement"];
	        this.idSyndique = source["idSyndique"];
	        this.numeroTelProf = source["numeroTelProf"];
	        this.premierAnCoti = source["premierAnCoti"];
	        this.syndique = source["syndique"];
	        this.titre = source["titre"];
	        this.typeInstallation = source["typeInstallation"];
	        this.categorieCotiCat = source["categorieCotiCat"];
	        this.categorieCotiDep = source["categorieCotiDep"];
	        this.adresseProf1 = source["adresseProf1"];
	        this.adresseProf2 = source["adresseProf2"];
	        this.age = source["age"];
	        this.association = source["association"];
	        this.codePostalProf = source["codePostalProf"];
	        this.dateDiplome = source["dateDiplome"];
	        this.dateModification = source["dateModification"];
	        this.dateNaissance = source["dateNaissance"];
	        this.diplomeFaculte = source["diplomeFaculte"];
	        this.lieuExercice = source["lieuExercice"];
	        this.numeroTelDomicile = source["numeroTelDomicile"];
	        this.paysProf = source["paysProf"];
	        this.persAideAssistante = source["persAideAssistante"];
	        this.persAssistante = source["persAssistante"];
	        this.persCollaborateur = source["persCollaborateur"];
	        this.persFemmeDeMenage = source["persFemmeDeMenage"];
	        this.persLaboratoire = source["persLaboratoire"];
	        this.persReceptionniste = source["persReceptionniste"];
	        this.personnels = source["personnels"];
	        this.remarques = source["remarques"];
	        this.responsable = source["responsable"];
	        this.sexe = source["sexe"];
	        this.typeExercice = source["typeExercice"];
	        this.villeProf = source["villeProf"];
	    }
	}

}

export namespace sql {
	
	export class DB {
	
	
	    static createFrom(source: any = {}) {
	        return new DB(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	
	    }
	}

}

