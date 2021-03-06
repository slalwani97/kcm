/*
Copyright 2018 The nullhash Authors.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package bootstrap

import (
	"log"
	"os"

	"github.com/nullhash/kcm/kcmmanager/config"
)

var (
	home string
)

func deleteHomeIfPresent() {
	if _, err := os.Stat(home + "/.kcm"); !os.IsNotExist(err) {
		os.RemoveAll(home + "/.kcm")
	}
}
func setupKCMHome() {
	os.Mkdir(home+"/.kcm", os.ModePerm)
}

func setupKCMConfig() {
	if file, err := os.Create(home + "/.kcm/config"); err == nil {
		defer file.Close()
	}
}

func Bootstrap() {
	home = os.Getenv("HOME")
	if home == "" {
		log.Println("error while reading environment variable")
		return
	}
	deleteHomeIfPresent()
	setupKCMHome()
	setupKCMConfig()
	defaultKubeConfigPath := home + "/.kube/config"
	config.LoadConfig(defaultKubeConfigPath)
}
