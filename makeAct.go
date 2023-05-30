/*Crée un fichier Act de base A.Villanueva*/
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var err error
	fichier := "act.txt"
	alms := 0 //Nombre d'alarmes créées
	cmds := 1 //Nombre de cmd créés

	space := 10 //Espace entre les alarmes et les commandes

	comentario := "/**/" //Récupère les commentaires du fichier de configuration cfg

	//Analysez s'il y a des arguments
	if len(os.Args) < 4 {
		fmt.Println("Erreur args \n Example : makeAct nom_fichier alms cmds space ")
		os.Exit(1)
	}

	fichier = os.Args[1] //Nom fichier

	alms, err = strconv.Atoi(os.Args[2]) //num. alms
	if err != nil {
		fmt.Println("Erreur num. alms :", err)
		os.Exit(1)
	}

	cmds, err = strconv.Atoi(os.Args[3]) //num. cmd
	if err != nil {
		fmt.Println("Erreur num. cmds :", err)
		os.Exit(1)
	}

	space, err = strconv.Atoi(os.Args[4]) //spacement
	if err != nil {
		fmt.Println("Erreur spacement :", err)
		os.Exit(1)
	}

	salidaFile, err := os.Create("act_" + fichier + ".txt") //fichier avec act_
	if err != nil {
		fmt.Println("Erreur lors de la création du fichier de sortie :", err)
		os.Exit(1)
	}
	defer salidaFile.Close()
	//Write num.alms
	writeNumAlms(salidaFile, alms) //Write Num Alms Virtuelle

	//écrire des ALMs
	for alm := space; alm <= alms*space; alm += space {
		writeAlm(salidaFile, comentario, alm)
	}

	//Write num. CMDs
	writeNumCmds(salidaFile, cmds) //Write Num Alms Virtuelle

	//écrire des commandes
	for cmd := space; cmd <= cmds*space; cmd += space {
		writeCmd(salidaFile, comentario, cmd)
	}

	//écrire la dernière partie act
	writeEnd(salidaFile)

	fmt.Println("Fichier de sortie créé avec succès.")

}

func writeNumAlms(salidaFile *os.File, valor int) {
	fmt.Fprintln(salidaFile, "NBR_ALM:", valor)
	fmt.Fprintln(salidaFile, "")
}

func writeAlm(salidaFile *os.File, comentario string, valor int) {
	fmt.Fprintln(salidaFile, comentario)
	fmt.Fprintln(salidaFile, "ALM_NUM:", valor)
	fmt.Fprintln(salidaFile, "ALM_OUT:", "0xF")
	fmt.Fprintln(salidaFile, "ALM_MASK:", "0")
	fmt.Fprintln(salidaFile, "ALM_CHG:", "0")
	fmt.Fprintln(salidaFile, "ALM_STATUS:", "0x11")

	fmt.Fprintln(salidaFile, "ALM_SCNEW:\tbegin")
	fmt.Fprintln(salidaFile, "\t\t\tinfo")
	fmt.Fprintln(salidaFile, "\t\tend")
	fmt.Fprintln(salidaFile, "ALM_SCOLD:\tbegin")
	fmt.Fprintln(salidaFile, "\t\t\tinfo")
	fmt.Fprintln(salidaFile, "\t\tend")
	fmt.Fprintln(salidaFile, "")
}

func writeNumCmds(salidaFile *os.File, valor int) {
	fmt.Fprintln(salidaFile, "/* COMMANDES */\n")
	fmt.Fprintln(salidaFile, "NBR_CMD:", valor)
	fmt.Fprintln(salidaFile, "")
}

func writeCmd(salidaFile *os.File, comentario string, valor int) {
	fmt.Fprintln(salidaFile, comentario)
	fmt.Fprintln(salidaFile, "CMD_NUM:", valor)
	fmt.Fprintln(salidaFile, "CMD_OUT:", "0xF")
	fmt.Fprintln(salidaFile, "CMD_STATUS:", "1")
	fmt.Fprintln(salidaFile, "CMD_SC:", "\tbegin")
	fmt.Fprintln(salidaFile, "\t\t\tinfo")
	fmt.Fprintln(salidaFile, "\t\tend")
	fmt.Fprintln(salidaFile, "")
}

func writeEnd(salidaFile *os.File) {

	fmt.Fprintln(salidaFile, "DATE_SC:", "\tbegin")
	fmt.Fprintln(salidaFile, "\t\t\tend\n")

	fmt.Fprintln(salidaFile, "COUNTER_SC:", "\tbegin")
	fmt.Fprintln(salidaFile, "\t\tend\n")

	fmt.Fprintln(salidaFile, "LOGIC_SC:", "\tbegin")
	fmt.Fprintln(salidaFile, "\t\tend\n")

	fmt.Fprintln(salidaFile, "CNTADD_SC:", "\tbegin")
	fmt.Fprintln(salidaFile, "\t\tend\n")

	fmt.Fprintln(salidaFile, "CNTADDF_SC:", "\tbegin")
	fmt.Fprintln(salidaFile, "\t\tend\n")

	fmt.Fprintln(salidaFile, "MASKCMD0:", "\t{")
	fmt.Fprintln(salidaFile, "\t\t}\n")

	fmt.Fprintln(salidaFile, "MASKCMD1:", "\t{")
	fmt.Fprintln(salidaFile, "\t\t}\n")

	fmt.Fprintln(salidaFile, "MASKCMD2:", "\t{")
	fmt.Fprintln(salidaFile, "\t\t}\n")

}
