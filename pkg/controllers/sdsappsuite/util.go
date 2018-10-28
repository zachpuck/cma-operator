package sdsappsuite

import (
	//"github.com/golang/glog"
	"github.com/juju/loggo"
	//"github.com/samsung-cnct/cma-operator/pkg/util/cmagrpc"
	//"github.com/spf13/viper"
	//"io/ioutil"
	//"k8s.io/apimachinery/pkg/apis/meta/v1"
	//"k8s.io/client-go/rest"
	//"k8s.io/client-go/tools/clientcmd"
	//"os"
	"github.com/samsung-cnct/cma-operator/pkg/util"
)

var (
	logger loggo.Logger
)

func (c *SDSAppSuiteController) SetLogger() {
	logger = util.GetModuleLogger("pkg.controllers.sdsappsuite", loggo.INFO)
}
