package controller

import (
	"database/sql"
	"fmt"
)

/*Metodo para realizar un query en una tabla específica, retorna un cursor que comienza en la primera
fila del resultado retornado*/
func Query(table string, email string, pass string, db *sql.DB) *sql.Rows {
	rows, err := db.Query("select * from " + table + " where EMAIL=" + email + " and PASSWORD=" + pass)
	if err != nil {
		fmt.Println("Error running query")
		fmt.Println(err)
	}
	return rows
}
func Connect() *sql.DB {
	db, err := sql.Open("godror", `user="system" password="root" connectString="localhost:1521/xe"
	parseTime=true`)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	return db
}

//func excecTour()
func Controller() {
	/*
		15:20
		Conección a la base de datos a través del paquete godror
	*/
	// db, err := sql.Open("godror", `user="system" password="root" connectString="localhost:1521/xe"
	// parseTime=true`)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// defer db.Close()
	// rows := query("USUARIO", email, pass, db)
	// defer rows.Close()
	// //var u1 usuario
	// var id int
	// var cod_empleado int
	// var fec_Ing sql.NullTime
	// var fec_Sal sql.NullTime
	// var nombre string
	// var salario int
	// var email string
	// var puesto int
	// var jefe string
	// var bonificacion string
	// for rows.Next() {
	// 	err := rows.Scan(&id, &cod_empleado, &fec_Ing, &fec_Sal, &nombre, &salario, &email, &puesto, &jefe, &bonificacion)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	fmt.Printf("ID:%v\ncod_empleado:%v\nfec_ing:%\nfec_salida:%s\nnombre:%s\n"+
	// 		"salario:%v\nemail:%s\npuesto:%v\njefe:%s\nbonificacion:%s\n",
	// 		id, cod_empleado, fec_Ing, fec_Sal, nombre, salario, email, puesto, jefe, bonificacion)
	// }

}
