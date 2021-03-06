package etcd

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/registry/generic"
	"k8s.io/apiserver/pkg/registry/generic/registry"
	"k8s.io/apiserver/pkg/registry/rest"
	"k8s.io/kubernetes/pkg/printers"
	printerstorage "k8s.io/kubernetes/pkg/printers/storage"

	networkapi "github.com/openshift/origin/pkg/network/apis/network"
	"github.com/openshift/origin/pkg/network/registry/netnamespace"
	printersinternal "github.com/openshift/origin/pkg/printers/internalversion"
	"github.com/openshift/origin/pkg/util/restoptions"
)

// rest implements a RESTStorage for sdn against etcd
type REST struct {
	*registry.Store
}

var _ rest.StandardStorage = &REST{}

// NewREST returns a RESTStorage object that will work against netnamespaces
func NewREST(optsGetter restoptions.Getter) (*REST, error) {
	store := &registry.Store{
		NewFunc:                  func() runtime.Object { return &networkapi.NetNamespace{} },
		NewListFunc:              func() runtime.Object { return &networkapi.NetNamespaceList{} },
		DefaultQualifiedResource: networkapi.Resource("netnamespaces"),

		TableConvertor: printerstorage.TableConvertor{TablePrinter: printers.NewTablePrinter().With(printersinternal.AddHandlers)},

		CreateStrategy: netnamespace.Strategy,
		UpdateStrategy: netnamespace.Strategy,
		DeleteStrategy: netnamespace.Strategy,
	}

	options := &generic.StoreOptions{RESTOptions: optsGetter}
	if err := store.CompleteWithOptions(options); err != nil {
		return nil, err
	}

	return &REST{store}, nil
}
