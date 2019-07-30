package main

import (
	"encoding/json"
	"log"
	"net/http"

)


func returnAllUsers(w http.ResponseWriter, r *http.Request) {
	var users News
	var arr_user []News
	var response Response

	db := connect()
	defer db.Close()

	rows, err := db.Query("Select a.person_id, author, channel, content from identity_info a inner join news b on a.person_id=b.person_id")
	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		if err := rows.Scan(&users.Id, &users.Author, &users.Channel, &users.Content, );
		err != nil {
			log.Fatal(err.Error())

		} else {
			arr_user = append(arr_user, users)
		}
	}

	response.Status = 1
	response.Message = "Success"
	response.Data = arr_user

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func addData(w http.ResponseWriter, r *http.Request) {

	var response Response
	var news News

	db := connect()
	defer db.Close()

	_=json.NewDecoder(r.Body).Decode(&news)

	_,err := db.Exec("insert into news (content, person_id) values(?,?)", news.Content, news.Person_id,)

	if err != nil {
		log.Fatal(err)
	}

	response.Status = 1
	response.Message = "Success Add"
	log.Print("Insert data to database")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

	
}

