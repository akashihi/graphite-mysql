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
	"fmt"
	"github.com/marpaia/graphite-golang"
)

func sendMetrics(status Status, config Configuration) {
	var Graphite, err = graphite.NewGraphite(config.MetricsHost, config.MetricsPort)
	if err != nil {
		log.Error("Can't connect to graphite collector: %v", err)
		return
	}
	Graphite.SimpleSend(fmt.Sprint(config.MetricsPrefix, ".mysql.connections"), status.Connections)
	Graphite.SimpleSend(fmt.Sprint(config.MetricsPrefix, ".mysql.bytesin"), status.BytesIn)
	Graphite.SimpleSend(fmt.Sprint(config.MetricsPrefix, ".mysql.bytesout"), status.BytesOut)
	Graphite.SimpleSend(fmt.Sprint(config.MetricsPrefix, ".mysql.queries"), status.Queries)
	Graphite.SimpleSend(fmt.Sprint(config.MetricsPrefix, ".mysql.threads.running"), status.ThreadsRunning)
}
