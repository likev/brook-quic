// Copyright (c) 2016-present Cloud <cloud@txthinking.com>
//
// This program is free software; you can redistribute it and/or
// modify it under the terms of version 3 of the GNU General Public
// License as published by the Free Software Foundation.
//
// This program is distributed in the hope that it will be useful, but
// WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU
// General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <https://www.gnu.org/licenses/>.

package main

import (
//	"context"
//	"crypto/x509"
//	"errors"
	"fmt"
//	"io/fs"
//	"io/ioutil"
//	"log"
//	"net"
	"os"
//	"os/exec"
//	"os/signal"
//	"path/filepath"
//	"runtime"
//	"strings"
//	"sync"
//	"syscall"
//	"time"

//	"net/http"
//	"net/url"

	"github.com/likev/brook"
//	"github.com/likev/brook/plugins/block"
//	"github.com/txthinking/brook/plugins/pprof"
//	"github.com/txthinking/brook/plugins/socks5dial"
//	"github.com/txthinking/brook/plugins/thedns"
//	"github.com/txthinking/brook/plugins/tproxy"
//	"github.com/txthinking/runnergroup"
//	"github.com/txthinking/socks5"
//	"github.com/urfave/cli/v2"
)

func main() {
	//g := runnergroup.New()
/* 	app := cli.NewApp()
	app.Name = "Brook"
	app.Version = "20230218"
	app.Usage = "A cross-platform network tool designed for developers"
	app.Authors = []*cli.Author{
		{
			Name:  "Cloud",
			Email: "cloud@txthinking.com",
		},
	}
	app.Copyright = "https://github.com/txthinking/brook"
	app.EnableBashCompletion = true
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:  "pprof",
			Usage: "go http pprof listen addr, such as :6060",
		},
	} */
	/*
	app.Before = func(c *cli.Context) error {
		if c.String("pprof") != "" {
			p, err := pprof.NewPprof(c.String("pprof"))
			if err != nil {
				return err
			}
			g.Add(&runnergroup.Runner{
				Start: func() error {
					return p.ListenAndServe()
				},
				Stop: func() error {
					return p.Shutdown()
				},
			})
		}
		return nil
	}
	*/

/*
	s, err := brook.NewWSServer(":" + os.Getenv("PORT"), os.Getenv("SECRET_WS"), "", "/ws", 0, 60, false)
	if err != nil {
		return
	}

	
*/

	s, err := brook.NewQUICServer("fly-global-services:" + os.Getenv("PORT_UDP"), os.Getenv("SECRET_WS"), os.Getenv("DOMAIN"), 0, 60, false)
	//s, err := brook.NewQUICServer(os.Getenv("PORT"), "someasffsd", "silent-brook-4924.fly.dev", 0, 60, false)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = s.ListenAndServe()
	if err != nil {
		fmt.Println(err)
		return
	}

	/*
	if os.Getenv("SOCKS5_DEBUG") != "" {
		socks5.Debug = true
	}
	*/

	//os.Args
/* 	osArgs := []string{"./brook", "wsserver"}
	if err := app.Run(osArgs); err != nil {
		log.Println(err)
		return
	} */
/* 	if len(g.Runners) == 0 {
		return
	} */

	/*
	go func() {
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
		<-sigs
		g.Done()
	}()
	*/
//	log.Println(g.Wait())
}