package config

import "fmt"

var logo = "      _           _    ____ ____ _____ \n  ___| |__   __ _| |_ / ___|  _ \\_   _|\n / __| '_ \\ / _` | __| |  _| |_) || |  \n| (__| | | | (_| | |_| |_| |  __/ | |  \n \\___|_| |_|\\__,_|\\__|\\____|_|    |_|     v: %s"

var Logo = fmt.Sprintf(logo, Version)