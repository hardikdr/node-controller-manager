package v1alpha1

import (
	v1alpha1 "github.com/gardener/node-controller-manager/pkg/apis/machine/v1alpha1"
	scheme "github.com/gardener/node-controller-manager/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// MachineDeploymentsGetter has a method to return a MachineDeploymentInterface.
// A group's client should implement this interface.
type MachineDeploymentsGetter interface {
	MachineDeployments(namespace string) MachineDeploymentInterface
}

// MachineDeploymentInterface has methods to work with MachineDeployment resources.
type MachineDeploymentInterface interface {
	Create(*v1alpha1.MachineDeployment) (*v1alpha1.MachineDeployment, error)
	Update(*v1alpha1.MachineDeployment) (*v1alpha1.MachineDeployment, error)
	UpdateStatus(*v1alpha1.MachineDeployment) (*v1alpha1.MachineDeployment, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.MachineDeployment, error)
	List(opts v1.ListOptions) (*v1alpha1.MachineDeploymentList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.MachineDeployment, err error)
	GetScale(machineDeploymentName string, options v1.GetOptions) (*v1alpha1.Scale, error)
	UpdateScale(machineDeploymentName string, scale *v1alpha1.Scale) (*v1alpha1.Scale, error)

	MachineDeploymentExpansion
}

// machineDeployments implements MachineDeploymentInterface
type machineDeployments struct {
	client rest.Interface
	ns     string
}

// newMachineDeployments returns a MachineDeployments
func newMachineDeployments(c *MachineV1alpha1Client, namespace string) *machineDeployments {
	return &machineDeployments{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the machineDeployment, and returns the corresponding machineDeployment object, and an error if there is any.
func (c *machineDeployments) Get(name string, options v1.GetOptions) (result *v1alpha1.MachineDeployment, err error) {
	result = &v1alpha1.MachineDeployment{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("machinedeployments").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of MachineDeployments that match those selectors.
func (c *machineDeployments) List(opts v1.ListOptions) (result *v1alpha1.MachineDeploymentList, err error) {
	result = &v1alpha1.MachineDeploymentList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("machinedeployments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested machineDeployments.
func (c *machineDeployments) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("machinedeployments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a machineDeployment and creates it.  Returns the server's representation of the machineDeployment, and an error, if there is any.
func (c *machineDeployments) Create(machineDeployment *v1alpha1.MachineDeployment) (result *v1alpha1.MachineDeployment, err error) {
	result = &v1alpha1.MachineDeployment{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("machinedeployments").
		Body(machineDeployment).
		Do().
		Into(result)
	return
}

// Update takes the representation of a machineDeployment and updates it. Returns the server's representation of the machineDeployment, and an error, if there is any.
func (c *machineDeployments) Update(machineDeployment *v1alpha1.MachineDeployment) (result *v1alpha1.MachineDeployment, err error) {
	result = &v1alpha1.MachineDeployment{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("machinedeployments").
		Name(machineDeployment.Name).
		Body(machineDeployment).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *machineDeployments) UpdateStatus(machineDeployment *v1alpha1.MachineDeployment) (result *v1alpha1.MachineDeployment, err error) {
	result = &v1alpha1.MachineDeployment{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("machinedeployments").
		Name(machineDeployment.Name).
		SubResource("status").
		Body(machineDeployment).
		Do().
		Into(result)
	return
}

// Delete takes name of the machineDeployment and deletes it. Returns an error if one occurs.
func (c *machineDeployments) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("machinedeployments").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *machineDeployments) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("machinedeployments").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched machineDeployment.
func (c *machineDeployments) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.MachineDeployment, err error) {
	result = &v1alpha1.MachineDeployment{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("machinedeployments").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}

// GetScale takes name of the machineDeployment, and returns the corresponding v1alpha1.Scale object, and an error if there is any.
func (c *machineDeployments) GetScale(machineDeploymentName string, options v1.GetOptions) (result *v1alpha1.Scale, err error) {
	result = &v1alpha1.Scale{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("machinedeployments").
		Name(machineDeploymentName).
		SubResource("scale").
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// UpdateScale takes the top resource name and the representation of a scale and updates it. Returns the server's representation of the scale, and an error, if there is any.
func (c *machineDeployments) UpdateScale(machineDeploymentName string, scale *v1alpha1.Scale) (result *v1alpha1.Scale, err error) {
	result = &v1alpha1.Scale{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("machinedeployments").
		Name(machineDeploymentName).
		SubResource("scale").
		Body(scale).
		Do().
		Into(result)
	return
}