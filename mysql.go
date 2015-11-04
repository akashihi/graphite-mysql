/*
   conntrack-logger
   Copyright (C) 2015 Denis V Chapligin <akashihi@gmail.com>
   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.
   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU General Public License for more details.
   You should have received a copy of the GNU General Public License
   along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	/*"bufio"
	  "fmt"
	  "net"
	  "strings"*/)

type Status struct {
	Connections    string
	BytesIn        string
	BytesOut       string
	Queries        string
	ThreadsRunning string
}

func getStatusData(host string, port int, user string, password string) (Status, error) {
	var result = Status{}

	var connectionUrl = fmt.Sprintf("%s:%s@tcp(%s:%d)/performance_schema?charset=utf8", user, password, host, port)

	log.Info("Connecting to %s", connectionUrl)
	db, err := sql.Open("mysql", connectionUrl)
	if err != nil {
		log.Error("Can't connect to mysql: %v", err)
		return Status{}, err
	}
	defer db.Close()

	rowsConnections, err := db.Query("SELECT COUNT(*) FROM information_schema.PROCESSLIST")
	if err != nil {
		log.Error("Can't retrieve process list: %v", err)
		return Status{}, err
	}
	defer rowsConnections.Close()

	for rowsConnections.Next() {
		err = rowsConnections.Scan(&result.Connections)
		if err != nil {
			log.Error("Can't fetch process list data: %v", err)
			return Status{}, err
		}
	}

	rowsGlobal, err := db.Query("SELECT * FROM information_schema.GLOBAL_STATUS where VARIABLE_NAME in ('BYTES_RECEIVED','BYTES_SENT','QUERIES','THREADS_RUNNING')")
	if err != nil {
		log.Error("Can't retrieve global variables: %v", err)
		return Status{}, err
	}
	defer rowsGlobal.Close()

	for rowsGlobal.Next() {
		var name string
		var value string
		err = rowsGlobal.Scan(&name, &value)
		if err != nil {
			log.Error("Can't fetch global variables data: %v", err)
			return Status{}, err
		}
		switch name {
		case "BYTES_RECEIVED":
			result.BytesIn = value
			break
		case "BYTES_SENT":
			result.BytesOut = value
			break
		case "QUERIES":
			result.Queries = value
			break
		case "THREADS_RUNNING":
			result.ThreadsRunning = value
			break
		}
	}

	return result, nil
}
